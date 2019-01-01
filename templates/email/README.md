# Email templates

## base.html

The base.html file mainly exists at the moment because I realize that I will
want to have some kind of standardization as time goes by and that even if I
only see one or two types of emails that I want to use for now, that may easily
change over time.

## invitation.html

The configuration has an InviteOnly flag. If that flag is true, then players
can only join the site if an administrator or someone with invite access sends
them an invitation.  The invitation will be an actual object in the database
that the user has to authenticate against, by following the invite link that
includes the invite token, in order to create an account.

## emainconfirmation.html

If a person is responding to an invitation and uses the same email address to
sign up, then it will be assumed that their address is verified. Otherwise, new
users will be asked to verify their email address.
