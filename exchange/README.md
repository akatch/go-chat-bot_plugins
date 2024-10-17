# Exchange

A plugin for running gift exchange campaigns in your channels.

## Dependencies

A `sqlite` database with the name $EXCHANGE_DB_NAME owned by $EXCHANGE_DB_ROLE

## Commands

exchange
  register -- subcommand for joining a Campaign
    address -- add freeform address
    notes -- add freeform notes
    show -- display this user's existing information

  campaign -- subcommands for managing Campaigns
    init -- initialize an empty campaign for a channel
    show -- display all registered users and their provided info
    shuffle -- assign each user a recipient. Users may not update notes or address after shuffle has occurred.
    end -- finish a campaign

## Environment Variables

EXCHANGE_ALLOWED_ADMINS="user1,user2,user3"
EXCHANGE_DB_NAME=exchange
EXCHANGE_DB_ROLE=exchange
EXCHANGE_DB_PASSWORD=exchange
