BroadCast can be implemented using channel, however, you need to implement a dispathcer since channel will only 
have one receiver receive the msg. Inside the dispathcer, you need to make multiple copy of the msgs and send out
to each reciver.

one example:
https://stackoverflow.com/questions/36417199/how-to-broadcast-message-using-channel