package constants

type HTTPStatusCode int

const HttpStatusCodes (
	Ok					HTTPStatusCode = 200
	Created				HTTPStatusCode = 201
	Accepted 			HTTPStatusCode = 202
	NotGound			HTTPStatusCode = 400
	Unauthorized 		HTTPStatusCode = 401
	Forbidden			HTTPStatusCode = 403
	BadRequest			HTTPStatusCode = 404
	TooManyRequests		HTTPStatusCode = 429
	TooManyRequests		HTTPStatusCode = 499
	InternalServerError	HTTPStatusCode = 500
	NotImplemented		HTTPStatusCode = 501
	badGateway			HTTPStatusCode = 502
	ServiceUnavailable	HTTPStatusCode = 503
	GatewayTimeout		HTTPStatusCode = 504
)

func (code HTTPStatusCode) String() string {
	switch code {
	case Ok:
		return "OK"
	case Created:
		return "Created"
	case Accepted:
		return "Accepted"
	case NoContent:
		return "No Content"
	case BadRequest:
		return "Bad Request"
	case Unauthorized:
		return "Unauthorized"
	case Forbidden:
		return "Forbidden"
	case NotFound:
		return "Not Found"
	case Conflict:
		return "Conflict"
	case TooManyRequests:
		return "Too many requests"
	case InternalServerError:
		return "Internal Server Error"
	case NotImplemented:
		return "Not Implemented"
	case BadGateway:
		return "Bad Gateway"
	case ServiceUnavailable:
		return "Service Unavailable"
	case GatewayTimeout:
		return "Gateway Timeout"
	default:
		return "Unknown Status"
	}
}