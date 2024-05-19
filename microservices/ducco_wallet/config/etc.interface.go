package config

type (
	EtcInterface struct {
		Orders Orders
	}

	Orders struct {
		OrdersStatus OrdersStatus
	}

	OrdersStatus struct {
		Init      uint8
		Success   uint8
		Delivered uint8
		Error     uint8
		Pending   uint8
	}
)
