//nolint
package config

// Bcrypt
const BCRYPT_WORK_FACTOR = 12
const BCRYPT_MAX_BYTES = 72

// Verification email
const EMAIL_VERIFICATION_TIMEOUT = TWELVE_HOURS

// sha1 -> 160 bits / 8 = 20 bytes * 2 (hex) = 40 bytes
const EMAIL_VERIFICATION_TOKEN_BYTES = 40

// sha256 -> 256 bits / 8 = 32 bytes * 2 (hex) = 64 bytes
const EMAIL_VERIFICATION_SIGNATURE_BYTES = 64

// Password reset
const PASSWORD_RESET_BYTES = 40
const PASSWORD_RESET_TIMEOUT = ONE_HOUR
