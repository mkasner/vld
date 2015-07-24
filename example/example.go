//go:generate vld -type=Transaction,Contract,Invoice -rp=git.aduro.hr/t2utils
package example

type Transaction struct {
	Hello string `vld:"req,maxlen=12"`
	World string `vld:"req,maxlen=12"`
}

type Contract struct {
	Price int64  `vld:"maxint64=20"`
	Type  string `vld:"req,maxlen=12"`
	Zero  string `vld:"req,maxlen=12"`
}

type Invoice struct {
	Amount   float32   `vld:"minfloat32=0"`
	Email    string    `vld:"email"`
	IP       string    `vld:"ip"`
	Status   *int      `vld:"req"`
	Contract *Contract `vld:"req"`
	MSISDN   string    `vld:"msisdn"`
}
