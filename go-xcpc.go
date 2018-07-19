// ======================================================================

//      cccccccccc          pppppppppp
//    cccccccccccccc      pppppppppppppp
//  ccccccccccccccc    ppppppppppppppppppp
// cccccc       cc    ppppppp        pppppp
// cccccc          pppppppp          pppppp
// cccccc        ccccpppp            pppppp
// cccccccc    cccccccc    pppp    ppppppp
//  ccccccccccccccccc     ppppppppppppppp
//     cccccccccccc      pppppppppppppp
//       cccccccc        pppppppppppp
//                       pppppp
//                       pppppp    

// ======================================================================
// Copyright © 2018. Counterparty Cash Association (CCA) Zug, CH.
// All Rights Reserved. All work owned by CCA is herby released 
// under Creative Commons Zero (0) License.

// Some rights of 3rd party, derivative and included works remain the
// property of thier respective owners. All marks, brands and logos of
// member groups remain the exclusive property of their owners and no
// right or endorsement is conferred by reference to thier organization 
// or brand(s) by CCA. 

// File:		go-bftx.go
// Description:	Applicaiton main file.

// Credits:

// Liming Jiang, Development
// Julian Smith, Direction
// Clayton Wong, Development
// Piotr Narewski, Gocoin Founder

// + Other contributors 

// =====================================================================

// Package xcpc manages Counterparty Cash (XCPC_TXs) transactions. As XCPC_TXs are executed
// or queried, the state is maintain in the local LevelDB databstore. Signed RAW transactions
// are parsed to gocoin-cash for transmission to the Bitcoin Cash blockchain.
package xcpc

import (

	// =======================
	// Golang Standard library
	// =======================
	"errors"	// Implements functions to manipulate errors
	"fmt"		// Implements formatted I/O with functions
	"io"		// Basic interfaces to I/O primitives
	"os"		// Platform-independent interface to operating system functionality
	"syscall"	// Interface to the low-level operating system primitives

	// ====================
	// Third-party packages
	// ====================
	"github.com/CounterpartyXCPC/gocoin-cash"	// Gocoin-cash is a full Bitcoin Cash solution written in Go language (golang)
	"github.com/syndtr/goleveldb/leveldb"		// Implementation of the LevelDB key/value database.
	
	// =========================
	// Conterparty Cash packages
	// =========================
	"github.com/CounterpartyXCPC/go-xcpc/encode"	// xcpc_tx encoder
)

var errNotSupported = errors.New("not supported")

// Run main application
func main() {
	// Do something
}