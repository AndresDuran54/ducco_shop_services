package config

type (
	EtcInterface struct {
		Sessions Sessions
	}

	Sessions struct {
		SessionsStatus SessionsStatus
	}

	SessionsStatus struct {
		Active   uint8
		Inactive uint8
	}
)
