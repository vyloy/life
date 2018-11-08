package gowasm

import . "syscall"

var errnoByCode = map[string]Errno{
	"EPERM":           EPERM,
	"ENOENT":          ENOENT,
	"ESRCH":           ESRCH,
	"EINTR":           EINTR,
	"EIO":             EIO,
	"ENXIO":           ENXIO,
	"E2BIG":           E2BIG,
	"ENOEXEC":         ENOEXEC,
	"EBADF":           EBADF,
	"ECHILD":          ECHILD,
	"EAGAIN":          EAGAIN,
	"ENOMEM":          ENOMEM,
	"EACCES":          EACCES,
	"EFAULT":          EFAULT,
	"EBUSY":           EBUSY,
	"EEXIST":          EEXIST,
	"EXDEV":           EXDEV,
	"ENODEV":          ENODEV,
	"ENOTDIR":         ENOTDIR,
	"EISDIR":          EISDIR,
	"EINVAL":          EINVAL,
	"ENFILE":          ENFILE,
	"EMFILE":          EMFILE,
	"ENOTTY":          ENOTTY,
	"EFBIG":           EFBIG,
	"ENOSPC":          ENOSPC,
	"ESPIPE":          ESPIPE,
	"EROFS":           EROFS,
	"EMLINK":          EMLINK,
	"EPIPE":           EPIPE,
	"ENAMETOOLONG":    ENAMETOOLONG,
	"ENOSYS":          ENOSYS,
	"EDQUOT":          EDQUOT,
	"EDOM":            EDOM,
	"ERANGE":          ERANGE,
	"EDEADLK":         EDEADLK,
	"ENOLCK":          ENOLCK,
	"ENOTEMPTY":       ENOTEMPTY,
	"ELOOP":           ELOOP,
	"ENOMSG":          ENOMSG,
	"EIDRM":           EIDRM,
	"ENOSTR":          ENOSTR,
	"ENODATA":         ENODATA,
	"ETIME":           ETIME,
	"ENOSR":           ENOSR,
	"EREMOTE":         EREMOTE,
	"ENOLINK":         ENOLINK,
	"EPROTO":          EPROTO,
	"EMULTIHOP":       EMULTIHOP,
	"EBADMSG":         EBADMSG,
	"EOVERFLOW":       EOVERFLOW,
	"EILSEQ":          EILSEQ,
	"EUSERS":          EUSERS,
	"ENOTSOCK":        ENOTSOCK,
	"EDESTADDRREQ":    EDESTADDRREQ,
	"EMSGSIZE":        EMSGSIZE,
	"EPROTOTYPE":      EPROTOTYPE,
	"ENOPROTOOPT":     ENOPROTOOPT,
	"EPROTONOSUPPORT": EPROTONOSUPPORT,
	"ESOCKTNOSUPPORT": ESOCKTNOSUPPORT,
	"EOPNOTSUPP":      EOPNOTSUPP,
	"EPFNOSUPPORT":    EPFNOSUPPORT,
	"EAFNOSUPPORT":    EAFNOSUPPORT,
	"EADDRINUSE":      EADDRINUSE,
	"EADDRNOTAVAIL":   EADDRNOTAVAIL,
	"ENETDOWN":        ENETDOWN,
	"ENETUNREACH":     ENETUNREACH,
	"ENETRESET":       ENETRESET,
	"ECONNABORTED":    ECONNABORTED,
	"ECONNRESET":      ECONNRESET,
	"ENOBUFS":         ENOBUFS,
	"EISCONN":         EISCONN,
	"ENOTCONN":        ENOTCONN,
	"ESHUTDOWN":       ESHUTDOWN,
	"ETOOMANYREFS":    ETOOMANYREFS,
	"ETIMEDOUT":       ETIMEDOUT,
	"ECONNREFUSED":    ECONNREFUSED,
	"EHOSTDOWN":       EHOSTDOWN,
	"EHOSTUNREACH":    EHOSTUNREACH,
	"EALREADY":        EALREADY,
	"EINPROGRESS":     EINPROGRESS,
	"ESTALE":          ESTALE,
	"ENOTSUP":         ENOTSUP,
	"ECANCELED":       ECANCELED,
	"EFTYPE":          EFTYPE,
	"EPROCLIM":        EPROCLIM,
	"EWOULDBLOCK":     EWOULDBLOCK,
}

var codeByErrno = map[Errno]string{
}

func init() {
	for key, value := range errnoByCode {
		codeByErrno[value] = key
	}
}