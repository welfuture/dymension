package types

import (
	"fmt"
	"net/url"

	errorsmod "cosmossdk.io/errors"
	"github.com/cockroachdb/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dymensionxyz/gerr-cosmos/gerrc"
)

// constant for maximum string length of the SequencerMetadata fields
const (
	MaxMonikerLength      = 70
	MaxContactFieldLength = 140
	MaxDetailsLength      = 280
	MaxExtraDataLength    = 280
	maxURLLength          = 256
	maxListLength         = 5
)

type AllowedDecimals uint32

const (
	Decimals18 AllowedDecimals = 18
)

func (d SequencerMetadata) Validate() error {
	_, err := d.EnsureLength()
	if err != nil {
		return err
	}

	if err = validateURLs(d.Rpcs); err != nil {
		return errorsmod.Wrap(err, "invalid rpcs URLs")
	}

	if err = validateURLs(d.RestApiUrls); err != nil {
		return errorsmod.Wrap(err, "invalid rest api URLs")
	}

	if d.FeeDenom != nil {
		if err := d.FeeDenom.Validate(); err != nil {
			return errors.Join(ErrInvalidFeeDenom, err)
		}
	}

	if d.ContactDetails == nil {
		return nil
	}

	return d.ContactDetails.Validate()
}

func (dm DenomMetadata) IsSet() bool {
	return dm != DenomMetadata{}
}

func (dm DenomMetadata) Validate() error {
	if err := sdk.ValidateDenom(dm.Base); err != nil {
		return fmt.Errorf("invalid metadata base denom: %w", err)
	}

	if err := sdk.ValidateDenom(dm.Display); err != nil {
		return fmt.Errorf("invalid metadata display denom: %w", err)
	}

	// validate exponent
	if AllowedDecimals(dm.Exponent) != Decimals18 {
		return fmt.Errorf("invalid exponent")
	}

	return nil
}

func (cd ContactDetails) Validate() error {
	if cd.Website != "" {
		if err := validateURL(cd.Website); err != nil {
			return errorsmod.Wrap(ErrInvalidURL, "invalid website URL")
		}
	}

	if cd.Telegram != "" {
		if err := validateURL(cd.Telegram); err != nil {
			return errorsmod.Wrap(ErrInvalidURL, "invalid telegram URL")
		}
	}

	if cd.X != "" {
		if err := validateURL(cd.X); err != nil {
			return errorsmod.Wrap(ErrInvalidURL, "invalid x URL")
		}
	}

	return nil
}

// ValidateURLs validates the URLs of a sequencer's metadata.
func validateURLs(urls []string) error {
	if len(urls) == 0 {
		return errorsmod.Wrap(gerrc.ErrInvalidArgument, "urls cannot be empty")
	}

	for _, u := range urls {
		if err := validateURL(u); err != nil {
			return errorsmod.Wrap(ErrInvalidURL, err.Error())
		}
	}

	return nil
}

func validateURL(urlStr string) error {
	if urlStr == "" {
		return errorsmod.Wrap(gerrc.ErrInvalidArgument, "url cannot be empty")
	}

	if len(urlStr) > maxURLLength {
		return fmt.Errorf("URL exceeds maximum length")
	}

	if _, err := url.ParseRequestURI(urlStr); err != nil {
		return fmt.Errorf("invalid URL: %w", err)
	}

	return nil
}

// EnsureLength ensures the length of a sequencer's metadata.
func (d SequencerMetadata) EnsureLength() (SequencerMetadata, error) {
	if len(d.Moniker) > MaxMonikerLength {
		return d, errors.Newf("invalid moniker length; got: %d, max: %d", len(d.Moniker), MaxMonikerLength)
	}

	if len(d.Details) > MaxDetailsLength {
		return d, errors.Newf("invalid details length; got: %d, max: %d", len(d.Details), MaxDetailsLength)
	}

	if len(d.ExtraData) > MaxExtraDataLength {
		return d, errors.Newf("invalid extra data length; got: %d, max: %d", len(d.ExtraData), MaxExtraDataLength)
	}

	if cd := d.ContactDetails; cd != nil {
		if len(cd.Website) > MaxContactFieldLength {
			return d, errors.Newf("invalid website length; got: %d, max: %d", len(cd.Website), MaxContactFieldLength)
		}
		if len(cd.Telegram) > MaxContactFieldLength {
			return d, errors.Newf("invalid telegram length; got: %d, max: %d", len(cd.Telegram), MaxContactFieldLength)
		}
		if len(cd.X) > MaxContactFieldLength {
			return d, errors.Newf("invalid x length; got: %d, max: %d", len(cd.X), MaxContactFieldLength)
		}
	}

	if len(d.P2PSeeds) > maxListLength {
		return d, errors.Newf("invalid p2p seeds length; got: %d, max: %d", len(d.P2PSeeds), maxListLength)
	}

	if len(d.Rpcs) > maxListLength {
		return d, errors.Newf("invalid rpcs length; got: %d, max: %d", len(d.Rpcs), maxListLength)
	}

	if len(d.EvmRpcs) > maxListLength {
		return d, errors.Newf("invalid evm rpcs length; got: %d, max: %d", len(d.EvmRpcs), maxListLength)
	}

	if len(d.RestApiUrls) > maxListLength {
		return d, errors.Newf("invalid rest api urls length; got: %d, max: %d", len(d.RestApiUrls), maxListLength)
	}

	if len(d.GenesisUrls) > maxListLength {
		return d, errors.Newf("invalid genesis urls length; got: %d, max: %d", len(d.GenesisUrls), maxListLength)
	}

	return d, nil
}
