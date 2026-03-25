package grpc

import (
    "fmt"
    "net/http"

    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

// ============================================================================
// GRPCError — обёртка над gRPC ошибкой
// ============================================================================

// GRPCError представляет ошибку gRPC вызова с дополнительным контекстом.
// Используется для единообразной обработки ошибок в слоях выше (service, handler).
type GRPCError struct {
    // Method — имя вызванного gRPC метода (например, "CreateOrder")
    Method string

    // Code — gRPC код ошибки (например, codes.NotFound, codes.Unavailable)
    Code codes.Code

    // Message — сообщение об ошибке от сервера
    Message string

    // Err — оригинальная ошибка для цепочки (Unwrap)
    Err error
}

// Error реализует интерфейс error для GRPCError.
// Возвращает человеко-читаемое представление ошибки.
//
// Пример вывода:
//   "gRPC CreateOrder failed [NOT_FOUND]: Order ord-123 not found"
func (e *GRPCError) Error() string {
    if e.Message != "" {
        return fmt.Sprintf("gRPC %s failed [%s]: %s", e.Method, e.Code, e.Message)
    }
    if e.Err != nil {
        return fmt.Sprintf("gRPC %s failed: %v", e.Method, e.Err)
    }
    return fmt.Sprintf("gRPC %s failed [%s]", e.Method, e.Code)
}

// Unwrap возвращает оригинальную ошибку.
// Позволяет использовать errors.Is() и errors.As() для проверки типа ошибки.
//
// Пример:
//   if errors.As(err, &grpcErr) { ... }
func (e *GRPCError) Unwrap() error {
    return e.Err
}

// ToHTTPStatus маппит gRPC код ошибки на соответствующий HTTP статус код.
// Используется в HTTP handler для возврата правильного статуса клиенту.
func (e *GRPCError) ToHTTPStatus() int {
    switch e.Code {
    case codes.OK:
        return http.StatusOK

    // Клиентские ошибки (4xx)
    case codes.Canceled:
        return http.StatusRequestTimeout // 408
    case codes.InvalidArgument:
        return http.StatusBadRequest // 400
    case codes.DeadlineExceeded:
        return http.StatusGatewayTimeout // 504
    case codes.NotFound:
        return http.StatusNotFound // 404
    case codes.AlreadyExists:
        return http.StatusConflict // 409
    case codes.PermissionDenied:
        return http.StatusForbidden // 403
    case codes.Unauthenticated:
        return http.StatusUnauthorized // 401
    case codes.ResourceExhausted:
        return http.StatusTooManyRequests // 429
    case codes.FailedPrecondition:
        return http.StatusPreconditionFailed // 412
    case codes.Aborted:
        return http.StatusConflict // 409
    case codes.OutOfRange:
        return http.StatusBadRequest // 400

    // Серверные ошибки (5xx)
    case codes.Unimplemented:
        return http.StatusNotImplemented // 501
    case codes.Internal:
        return http.StatusInternalServerError // 500
    case codes.Unavailable:
        return http.StatusServiceUnavailable // 503
    case codes.DataLoss:
        return http.StatusInternalServerError // 500

    // По умолчанию — внутренняя ошибка сервера
    default:
        return http.StatusInternalServerError
    }
}

// IsCode проверяет, соответствует ли ошибка указанному gRPC коду.
// Удобная альтернатива ручному приведению типа и проверке .Code.
func (e *GRPCError) IsCode(code codes.Code) bool {
    return e.Code == code
}

// NewGRPCError создаёт новую обёртку GRPCError из gRPC ошибки.
//
// Параметры:
//   • err — оригинальная ошибка от gRPC вызова
//   • method — имя вызванного метода для контекста (например, "CreateOrder")
//
// Возвращает:
//   • *GRPCError — обёрнутую ошибку с извлечённым status.Code и status.Message
func NewGRPCError(err error, method string) *GRPCError {
    if err == nil {
        return nil
    }

    grpcErr := &GRPCError{
        Method: method,
        Err:    err,
    }

    if st, ok := status.FromError(err); ok {
        grpcErr.Code = st.Code()
        grpcErr.Message = st.Message()
    } else {
        grpcErr.Code = codes.Unknown
        grpcErr.Message = err.Error()
    }

    return grpcErr
}

// NewGRPCErrorWithMessage создаёт GRPCError с переопределённым сообщением.
// Полезно когда нужно скрыть или переформатировать сообщение от сервера.
func NewGRPCErrorWithMessage(err error, method, message string) *GRPCError {
    grpcErr := NewGRPCError(err, method)
    if grpcErr != nil {
        grpcErr.Message = message
    }
    return grpcErr
}

// IsGRPCError проверяет, является ли ошибка обёрткой GRPCError.
// Удобная альтернатива ручному errors.As().
func IsGRPCError(err error) bool {
    var grpcErr *GRPCError
    return err != nil && AsGRPCError(err, &grpcErr)
}

// AsGRPCError пытается привести ошибку к типу *GRPCError.
// Аналог errors.As() но с более понятным именем для этого конкретного типа.
func AsGRPCError(err error, target **GRPCError) bool {
    if err == nil || target == nil {
        return false
    }
    return errAsGRPCError(err, target)
}

// Внутренняя реализация для AsGRPCError
func errAsGRPCError(err error, target **GRPCError) bool {
    if grpcErr, ok := err.(*GRPCError); ok {
        *target = grpcErr
        return true
    }
    if unwrapper, ok := err.(interface{ Unwrap() error }); ok {
        return errAsGRPCError(unwrapper.Unwrap(), target)
    }
    return false
}

// IsNotFound проверяет, является ли ошибка кодом codes.NotFound.
// Удобно для проверки "ресурс не найден" без ручного приведения типа.
func IsNotFound(err error) bool {
    var grpcErr *GRPCError
    return AsGRPCError(err, &grpcErr) && grpcErr.Code == codes.NotFound
}

// IsUnavailable проверяет, является ли ошибка кодом codes.Unavailable.
// Полезно для реализации retry-логики при временной недоступности сервиса.
func IsUnavailable(err error) bool {
    var grpcErr *GRPCError
    return AsGRPCError(err, &grpcErr) && grpcErr.Code == codes.Unavailable
}

// IsDeadlineExceeded проверяет, является ли ошибка кодом codes.DeadlineExceeded.
// Полезно для различения таймаута сети от бизнес-ошибок.
func IsDeadlineExceeded(err error) bool {
    var grpcErr *GRPCError
    return AsGRPCError(err, &grpcErr) && grpcErr.Code == codes.DeadlineExceeded
}

// IsInvalidArgument проверяет, является ли ошибка кодом codes.InvalidArgument.
// Полезно для валидации входных данных на стороне клиента.
func IsInvalidArgument(err error) bool {
    var grpcErr *GRPCError
    return AsGRPCError(err, &grpcErr) && grpcErr.Code == codes.InvalidArgument
}

// IsAlreadyExists проверяет, является ли ошибка кодом codes.AlreadyExists.
// Полезно для обработки конфликтов при создании ресурсов.
func IsAlreadyExists(err error) bool {
    var grpcErr *GRPCError
    return AsGRPCError(err, &grpcErr) && grpcErr.Code == codes.AlreadyExists
}

// IsPermissionDenied проверяет, является ли ошибка кодом codes.PermissionDenied.
// Полезно для обработки ошибок авторизации.
func IsPermissionDenied(err error) bool {
    var grpcErr *GRPCError
    return AsGRPCError(err, &grpcErr) && grpcErr.Code == codes.PermissionDenied
}

// IsUnauthenticated проверяет, является ли ошибка кодом codes.Unauthenticated.
// Полезно для обработки ошибок аутентификации.
func IsUnauthenticated(err error) bool {
    var grpcErr *GRPCError
    return AsGRPCError(err, &grpcErr) && grpcErr.Code == codes.Unauthenticated
}

// IsResourceExhausted проверяет, является ли ошибка кодом codes.ResourceExhausted.
// Полезно для обработки rate limiting и квот.
func IsResourceExhausted(err error) bool {
    var grpcErr *GRPCError
    return AsGRPCError(err, &grpcErr) && grpcErr.Code == codes.ResourceExhausted
}

// IsInternal проверяет, является ли ошибка кодом codes.Internal.
// Полезно для логирования серверных ошибок.
func IsInternal(err error) bool {
    var grpcErr *GRPCError
    return AsGRPCError(err, &grpcErr) && grpcErr.Code == codes.Internal
}

// GRPCErrorCode возвращает gRPC код ошибки если она является *GRPCError.
// Если нет — возвращает codes.Unknown.
func GRPCErrorCode(err error) codes.Code {
    var grpcErr *GRPCError
    if AsGRPCError(err, &grpcErr) {
        return grpcErr.Code
    }
    return codes.Unknown
}

// GRPCErrorMessage возвращает сообщение об ошибке если она является *GRPCError.
// Если нет — возвращает err.Error() или пустую строку.
func GRPCErrorMessage(err error) string {
    var grpcErr *GRPCError
    if AsGRPCError(err, &grpcErr) {
        return grpcErr.Message
    }
    if err != nil {
        return err.Error()
    }
    return ""
}

// GRPCErrorDetails возвращает детали ошибки в удобном для логирования формате.
func GRPCErrorDetails(err error) map[string]string {
    details := make(map[string]string)

    var grpcErr *GRPCError
    if AsGRPCError(err, &grpcErr) {
        details["method"] = grpcErr.Method
        details["code"] = grpcErr.Code.String()
        details["message"] = grpcErr.Message
        if grpcErr.Err != nil && grpcErr.Err.Error() != grpcErr.Message {
            details["original"] = grpcErr.Err.Error()
        }
        return details
    }

    if err != nil {
        details["error"] = err.Error()
    }
    return details
}