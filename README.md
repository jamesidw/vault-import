# Bring Your Own Key

Adding a key to a HashiCorp Vault demands that you do some AES encryption first.

This is a fleshed out implementation of the example they provide in Go

The general idea is that you want to wrap the actual key using AES then send the AES key to vault as an encrypted string (basic RSA encryption).
Since vault knows the private key, it can decrypt the key and unwrap the material you want to import

Read the original instructions from the [key wrapping guide](https://developer.hashicorp.com/vault/docs/secrets/transit/key-wrapping-guide)