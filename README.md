Monkey IDP
==========

This is me monkeying around.  I'm building an identity provider to see if I can.  When
it works, I'm going to get feedback on it to make it better.

## Current Features

### Language Choices

* API is in Golang because that's how it should be.
* UI is in ReactJS because it works well.
* DB is in PostgreSQL because there's no other choice in databases.

### Security

* There are scripts to create a root certificate authority and to generate a site cert for web traffic.  The purpose of this feature is to encrypt authentication requests while in-transit.
* The API users bcrypt.GenerateFromPassword() to store the salted hash of the user's password.  That's the best method I could find for storing passwords, for now. The purpose of this feature is to ensure that passwords are never stored in a manner that can be easily reverse engineered.

__*Question*__:  Should we encrypt it also?
