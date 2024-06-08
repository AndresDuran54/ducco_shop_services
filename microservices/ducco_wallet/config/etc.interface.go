package config

type (
	EtcInterface struct {
		Orders     Orders
		Parameters Parameters
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

	Parameters struct {
		ParametersTypes ParametersTypes
	}

	ParametersTypes struct {
		Number uint8
		String uint8
		JSON   uint8
	}
)
