package dcrm 

import (
	"math/big"
	"github.com/ethereum/go-ethereum/crypto/pbc"
)

type CmtMasterPublicKey struct {
	g *pbc.Element
	q *big.Int
	h *pbc.Element
	pairing *pbc.Pairing
}

func (cmpk *CmtMasterPublicKey) New(g *pbc.Element,q *big.Int,h *pbc.Element,pairing *pbc.Pairing) {
    cmpk.g = g
    cmpk.q = q
    cmpk.h = h
    cmpk.pairing = pairing
}
