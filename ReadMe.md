# Simple Socket Layer4 Loadbalancer written in GO.

In this learning project I create a Proxy with a simple roundroubin as Load Balancing.
The Idea is learning by build basic IT Components new on a very basic level.

## To Dos:
Open:
- Add Unit Tests
- Add Filter that after one new User crete a request any time get the same target
- Build really HA Postgres Example for better testing.
- How to run the LoadBalancer as HA?
- Add better Error Handling

Bugs:
- The server crashes than at start time one of the target endpoints is not exist.


What I have learned ?
- How to open Network Socket Server
- How to send Data to Socket Pars
- How to copy complete Stream with io.copy and send to target server.
- How setup a basic roundroubin (any new connection get the next target). Still in process.
