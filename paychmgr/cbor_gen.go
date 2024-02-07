// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package paychmgr

import (
	"fmt"
	"io"
	"math"
	"sort"

	cid "github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"

	address "github.com/filecoin-project/go-address"
	paych "github.com/filecoin-project/go-state-types/builtin/v8/paych"
)

var _ = xerrors.Errorf
var _ = cid.Undef
var _ = math.E
var _ = sort.Sort

func (t *VoucherInfo) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write([]byte{163}); err != nil {
		return err
	}

	// t.Proof ([]uint8) (slice)
	if len("Proof") > 8192 {
		return xerrors.Errorf("Value in field \"Proof\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("Proof"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("Proof")); err != nil {
		return err
	}

	if len(t.Proof) > 2097152 {
		return xerrors.Errorf("Byte array in field t.Proof was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajByteString, uint64(len(t.Proof))); err != nil {
		return err
	}

	if _, err := cw.Write(t.Proof); err != nil {
		return err
	}

	// t.Voucher (paych.SignedVoucher) (struct)
	if len("Voucher") > 8192 {
		return xerrors.Errorf("Value in field \"Voucher\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("Voucher"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("Voucher")); err != nil {
		return err
	}

	if err := t.Voucher.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.Submitted (bool) (bool)
	if len("Submitted") > 8192 {
		return xerrors.Errorf("Value in field \"Submitted\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("Submitted"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("Submitted")); err != nil {
		return err
	}

	if err := cbg.WriteBool(w, t.Submitted); err != nil {
		return err
	}
	return nil
}

func (t *VoucherInfo) UnmarshalCBOR(r io.Reader) (err error) {
	*t = VoucherInfo{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("VoucherInfo: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadStringWithMax(cr, 8192)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.Proof ([]uint8) (slice)
		case "Proof":

			maj, extra, err = cr.ReadHeader()
			if err != nil {
				return err
			}

			if extra > 2097152 {
				return fmt.Errorf("t.Proof: byte array too large (%d)", extra)
			}
			if maj != cbg.MajByteString {
				return fmt.Errorf("expected byte array")
			}

			if extra > 0 {
				t.Proof = make([]uint8, extra)
			}

			if _, err := io.ReadFull(cr, t.Proof); err != nil {
				return err
			}

			// t.Voucher (paych.SignedVoucher) (struct)
		case "Voucher":

			{

				b, err := cr.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := cr.UnreadByte(); err != nil {
						return err
					}
					t.Voucher = new(paych.SignedVoucher)
					if err := t.Voucher.UnmarshalCBOR(cr); err != nil {
						return xerrors.Errorf("unmarshaling t.Voucher pointer: %w", err)
					}
				}

			}
			// t.Submitted (bool) (bool)
		case "Submitted":

			maj, extra, err = cr.ReadHeader()
			if err != nil {
				return err
			}
			if maj != cbg.MajOther {
				return fmt.Errorf("booleans must be major type 7")
			}
			switch extra {
			case 20:
				t.Submitted = false
			case 21:
				t.Submitted = true
			default:
				return fmt.Errorf("booleans are either major type 7, value 20 or 21 (got %d)", extra)
			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
func (t *ChannelInfo) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write([]byte{174}); err != nil {
		return err
	}

	// t.Amount (big.Int) (struct)
	if len("Amount") > 8192 {
		return xerrors.Errorf("Value in field \"Amount\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("Amount"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("Amount")); err != nil {
		return err
	}

	if err := t.Amount.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.Target (address.Address) (struct)
	if len("Target") > 8192 {
		return xerrors.Errorf("Value in field \"Target\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("Target"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("Target")); err != nil {
		return err
	}

	if err := t.Target.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.Channel (address.Address) (struct)
	if len("Channel") > 8192 {
		return xerrors.Errorf("Value in field \"Channel\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("Channel"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("Channel")); err != nil {
		return err
	}

	if err := t.Channel.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.Control (address.Address) (struct)
	if len("Control") > 8192 {
		return xerrors.Errorf("Value in field \"Control\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("Control"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("Control")); err != nil {
		return err
	}

	if err := t.Control.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.NextLane (uint64) (uint64)
	if len("NextLane") > 8192 {
		return xerrors.Errorf("Value in field \"NextLane\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("NextLane"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("NextLane")); err != nil {
		return err
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.NextLane)); err != nil {
		return err
	}

	// t.Settling (bool) (bool)
	if len("Settling") > 8192 {
		return xerrors.Errorf("Value in field \"Settling\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("Settling"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("Settling")); err != nil {
		return err
	}

	if err := cbg.WriteBool(w, t.Settling); err != nil {
		return err
	}

	// t.Vouchers ([]*paychmgr.VoucherInfo) (slice)
	if len("Vouchers") > 8192 {
		return xerrors.Errorf("Value in field \"Vouchers\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("Vouchers"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("Vouchers")); err != nil {
		return err
	}

	if len(t.Vouchers) > 8192 {
		return xerrors.Errorf("Slice value in field t.Vouchers was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajArray, uint64(len(t.Vouchers))); err != nil {
		return err
	}
	for _, v := range t.Vouchers {
		if err := v.MarshalCBOR(cw); err != nil {
			return err
		}

	}

	// t.ChannelID (string) (string)
	if len("ChannelID") > 8192 {
		return xerrors.Errorf("Value in field \"ChannelID\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("ChannelID"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("ChannelID")); err != nil {
		return err
	}

	if len(t.ChannelID) > 8192 {
		return xerrors.Errorf("Value in field t.ChannelID was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(t.ChannelID))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string(t.ChannelID)); err != nil {
		return err
	}

	// t.CreateMsg (cid.Cid) (struct)
	if len("CreateMsg") > 8192 {
		return xerrors.Errorf("Value in field \"CreateMsg\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("CreateMsg"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("CreateMsg")); err != nil {
		return err
	}

	if t.CreateMsg == nil {
		if _, err := cw.Write(cbg.CborNull); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteCid(cw, *t.CreateMsg); err != nil {
			return xerrors.Errorf("failed to write cid field t.CreateMsg: %w", err)
		}
	}

	// t.Direction (uint64) (uint64)
	if len("Direction") > 8192 {
		return xerrors.Errorf("Value in field \"Direction\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("Direction"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("Direction")); err != nil {
		return err
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.Direction)); err != nil {
		return err
	}

	// t.AddFundsMsg (cid.Cid) (struct)
	if len("AddFundsMsg") > 8192 {
		return xerrors.Errorf("Value in field \"AddFundsMsg\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("AddFundsMsg"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("AddFundsMsg")); err != nil {
		return err
	}

	if t.AddFundsMsg == nil {
		if _, err := cw.Write(cbg.CborNull); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteCid(cw, *t.AddFundsMsg); err != nil {
			return xerrors.Errorf("failed to write cid field t.AddFundsMsg: %w", err)
		}
	}

	// t.PendingAmount (big.Int) (struct)
	if len("PendingAmount") > 8192 {
		return xerrors.Errorf("Value in field \"PendingAmount\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("PendingAmount"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("PendingAmount")); err != nil {
		return err
	}

	if err := t.PendingAmount.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.AvailableAmount (big.Int) (struct)
	if len("AvailableAmount") > 8192 {
		return xerrors.Errorf("Value in field \"AvailableAmount\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("AvailableAmount"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("AvailableAmount")); err != nil {
		return err
	}

	if err := t.AvailableAmount.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.PendingAvailableAmount (big.Int) (struct)
	if len("PendingAvailableAmount") > 8192 {
		return xerrors.Errorf("Value in field \"PendingAvailableAmount\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("PendingAvailableAmount"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("PendingAvailableAmount")); err != nil {
		return err
	}

	if err := t.PendingAvailableAmount.MarshalCBOR(cw); err != nil {
		return err
	}
	return nil
}

func (t *ChannelInfo) UnmarshalCBOR(r io.Reader) (err error) {
	*t = ChannelInfo{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("ChannelInfo: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadStringWithMax(cr, 8192)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.Amount (big.Int) (struct)
		case "Amount":

			{

				if err := t.Amount.UnmarshalCBOR(cr); err != nil {
					return xerrors.Errorf("unmarshaling t.Amount: %w", err)
				}

			}
			// t.Target (address.Address) (struct)
		case "Target":

			{

				if err := t.Target.UnmarshalCBOR(cr); err != nil {
					return xerrors.Errorf("unmarshaling t.Target: %w", err)
				}

			}
			// t.Channel (address.Address) (struct)
		case "Channel":

			{

				b, err := cr.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := cr.UnreadByte(); err != nil {
						return err
					}
					t.Channel = new(address.Address)
					if err := t.Channel.UnmarshalCBOR(cr); err != nil {
						return xerrors.Errorf("unmarshaling t.Channel pointer: %w", err)
					}
				}

			}
			// t.Control (address.Address) (struct)
		case "Control":

			{

				if err := t.Control.UnmarshalCBOR(cr); err != nil {
					return xerrors.Errorf("unmarshaling t.Control: %w", err)
				}

			}
			// t.NextLane (uint64) (uint64)
		case "NextLane":

			{

				maj, extra, err = cr.ReadHeader()
				if err != nil {
					return err
				}
				if maj != cbg.MajUnsignedInt {
					return fmt.Errorf("wrong type for uint64 field")
				}
				t.NextLane = uint64(extra)

			}
			// t.Settling (bool) (bool)
		case "Settling":

			maj, extra, err = cr.ReadHeader()
			if err != nil {
				return err
			}
			if maj != cbg.MajOther {
				return fmt.Errorf("booleans must be major type 7")
			}
			switch extra {
			case 20:
				t.Settling = false
			case 21:
				t.Settling = true
			default:
				return fmt.Errorf("booleans are either major type 7, value 20 or 21 (got %d)", extra)
			}
			// t.Vouchers ([]*paychmgr.VoucherInfo) (slice)
		case "Vouchers":

			maj, extra, err = cr.ReadHeader()
			if err != nil {
				return err
			}

			if extra > 8192 {
				return fmt.Errorf("t.Vouchers: array too large (%d)", extra)
			}

			if maj != cbg.MajArray {
				return fmt.Errorf("expected cbor array")
			}

			if extra > 0 {
				t.Vouchers = make([]*VoucherInfo, extra)
			}

			for i := 0; i < int(extra); i++ {
				{
					var maj byte
					var extra uint64
					var err error
					_ = maj
					_ = extra
					_ = err

					{

						b, err := cr.ReadByte()
						if err != nil {
							return err
						}
						if b != cbg.CborNull[0] {
							if err := cr.UnreadByte(); err != nil {
								return err
							}
							t.Vouchers[i] = new(VoucherInfo)
							if err := t.Vouchers[i].UnmarshalCBOR(cr); err != nil {
								return xerrors.Errorf("unmarshaling t.Vouchers[i] pointer: %w", err)
							}
						}

					}

				}
			}
			// t.ChannelID (string) (string)
		case "ChannelID":

			{
				sval, err := cbg.ReadStringWithMax(cr, 8192)
				if err != nil {
					return err
				}

				t.ChannelID = string(sval)
			}
			// t.CreateMsg (cid.Cid) (struct)
		case "CreateMsg":

			{

				b, err := cr.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := cr.UnreadByte(); err != nil {
						return err
					}

					c, err := cbg.ReadCid(cr)
					if err != nil {
						return xerrors.Errorf("failed to read cid field t.CreateMsg: %w", err)
					}

					t.CreateMsg = &c
				}

			}
			// t.Direction (uint64) (uint64)
		case "Direction":

			{

				maj, extra, err = cr.ReadHeader()
				if err != nil {
					return err
				}
				if maj != cbg.MajUnsignedInt {
					return fmt.Errorf("wrong type for uint64 field")
				}
				t.Direction = uint64(extra)

			}
			// t.AddFundsMsg (cid.Cid) (struct)
		case "AddFundsMsg":

			{

				b, err := cr.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := cr.UnreadByte(); err != nil {
						return err
					}

					c, err := cbg.ReadCid(cr)
					if err != nil {
						return xerrors.Errorf("failed to read cid field t.AddFundsMsg: %w", err)
					}

					t.AddFundsMsg = &c
				}

			}
			// t.PendingAmount (big.Int) (struct)
		case "PendingAmount":

			{

				if err := t.PendingAmount.UnmarshalCBOR(cr); err != nil {
					return xerrors.Errorf("unmarshaling t.PendingAmount: %w", err)
				}

			}
			// t.AvailableAmount (big.Int) (struct)
		case "AvailableAmount":

			{

				if err := t.AvailableAmount.UnmarshalCBOR(cr); err != nil {
					return xerrors.Errorf("unmarshaling t.AvailableAmount: %w", err)
				}

			}
			// t.PendingAvailableAmount (big.Int) (struct)
		case "PendingAvailableAmount":

			{

				if err := t.PendingAvailableAmount.UnmarshalCBOR(cr); err != nil {
					return xerrors.Errorf("unmarshaling t.PendingAvailableAmount: %w", err)
				}

			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
func (t *MsgInfo) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write([]byte{164}); err != nil {
		return err
	}

	// t.Err (string) (string)
	if len("Err") > 8192 {
		return xerrors.Errorf("Value in field \"Err\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("Err"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("Err")); err != nil {
		return err
	}

	if len(t.Err) > 8192 {
		return xerrors.Errorf("Value in field t.Err was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(t.Err))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string(t.Err)); err != nil {
		return err
	}

	// t.MsgCid (cid.Cid) (struct)
	if len("MsgCid") > 8192 {
		return xerrors.Errorf("Value in field \"MsgCid\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("MsgCid"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("MsgCid")); err != nil {
		return err
	}

	if err := cbg.WriteCid(cw, t.MsgCid); err != nil {
		return xerrors.Errorf("failed to write cid field t.MsgCid: %w", err)
	}

	// t.Received (bool) (bool)
	if len("Received") > 8192 {
		return xerrors.Errorf("Value in field \"Received\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("Received"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("Received")); err != nil {
		return err
	}

	if err := cbg.WriteBool(w, t.Received); err != nil {
		return err
	}

	// t.ChannelID (string) (string)
	if len("ChannelID") > 8192 {
		return xerrors.Errorf("Value in field \"ChannelID\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("ChannelID"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("ChannelID")); err != nil {
		return err
	}

	if len(t.ChannelID) > 8192 {
		return xerrors.Errorf("Value in field t.ChannelID was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(t.ChannelID))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string(t.ChannelID)); err != nil {
		return err
	}
	return nil
}

func (t *MsgInfo) UnmarshalCBOR(r io.Reader) (err error) {
	*t = MsgInfo{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("MsgInfo: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadStringWithMax(cr, 8192)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.Err (string) (string)
		case "Err":

			{
				sval, err := cbg.ReadStringWithMax(cr, 8192)
				if err != nil {
					return err
				}

				t.Err = string(sval)
			}
			// t.MsgCid (cid.Cid) (struct)
		case "MsgCid":

			{

				c, err := cbg.ReadCid(cr)
				if err != nil {
					return xerrors.Errorf("failed to read cid field t.MsgCid: %w", err)
				}

				t.MsgCid = c

			}
			// t.Received (bool) (bool)
		case "Received":

			maj, extra, err = cr.ReadHeader()
			if err != nil {
				return err
			}
			if maj != cbg.MajOther {
				return fmt.Errorf("booleans must be major type 7")
			}
			switch extra {
			case 20:
				t.Received = false
			case 21:
				t.Received = true
			default:
				return fmt.Errorf("booleans are either major type 7, value 20 or 21 (got %d)", extra)
			}
			// t.ChannelID (string) (string)
		case "ChannelID":

			{
				sval, err := cbg.ReadStringWithMax(cr, 8192)
				if err != nil {
					return err
				}

				t.ChannelID = string(sval)
			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
