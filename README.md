# safehouse-tech-back
Backend microservice for safehouse personal website

# Tech stack

# Endpoints



Critical Issues (High Priority)


4. Insufficient Input Validation - src/internal/middleware/http_validation.go:10-16
   - Only validates URL length (1000 chars) and HTTP methods
   - No validation of query parameters, headers, or request body content
5. Overly Permissive CSP - src/internal/middleware/security_headers.go:24
   - Development CSP allows 'unsafe-inline' and 'unsafe-eval'
   - Should be restricted even in development
6. Information Disclosure - src/internal/database/errors/database_errors.go:25-35
   - Generic error messages good, but internal errors could leak info
   - Consider adding request ID for debugging
