package rpcArg

type Arith int

type Args struct {
    A, B int
}

type Quotient struct {
    Quo, Rem int
}

func (t *Arith) Multiply(args *Args, reply *int) error {
    *reply = args.A * args.B
    return nil
}
