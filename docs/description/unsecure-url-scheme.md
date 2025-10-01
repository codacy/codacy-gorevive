## unsecure-url-scheme

_Description_: Checks for usage of potentially unsecure URL schemes (`http`, `ws`) in string literals.
Using unencrypted URL schemes can expose sensitive data during transmission and
make applications vulnerable to man-in-the-middle attacks.
Secure alternatives like `https` should be preferred when possible.

_Configuration_: N/A

The rule will not warn on local URLs (`localhost`, `127.0.0.1`).

