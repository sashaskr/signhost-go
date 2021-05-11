package signhost

const (
	TransactionWaitingForDocument = 5
	TransactionWaitingForSigner   = iota * 10
	TransactionInProgress
	TransactionSigned
	TransactionRejected
	TransactionExpired
	TransactionCanceled
	TransactionFailed
)

const (
	ActivityInvitationSent = iota + 101
	ActivityReceived
	ActivityOpened
	ActivityReminderSent
	ActivityDocumentOpened
)

const (
	ActivityCancelled = iota + 201
	ActivityRejected
	ActivitySigned
)

const (
	ActivitySignedDocumentSent = iota + 301
	ActivitySignedDocumentOpened
	ActivitySignedDocumentDownloaded
)

const (
	ActivityReceiptSent = iota + 401
	ActivityReceiptOpened
	ActivityReceiptDownloaded
)

const (
	ActivityFinished = 100 * (iota + 5)
	ActivityDeleted
	ActivityExpired
)

const (
	ActivityFailed = 999
)
