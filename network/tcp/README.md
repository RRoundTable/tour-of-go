# Reliable TCP Datastream

## Transmission Control Protocol

- Ordered data tranfer: the destination host rearranges segments according to sequence number.
- Retransmission of lost packets: any cumulative stream not acknowledged is retransmitted.
- Error-free data transfer: corrupted packtets are treated as lost and are retransmitted.
- Flow control: limits the rate a sender transfer data to gurantee reliable delivery. The receiver continually hints the sender on how much data can be received. When the receiving host's buffer fills, the next acknowlegment suspends the transfer and allows the data in the buffer to be processed.
- Congestion control: lost packets trigger a reduction in data delivery rate.


## TCP Session with 3-Way Handshake


1. Active open: The client Sends `SYN` to the server. The client sets the segment's sequence number to a random value A.
2. The server replies with a `SYN-ACK`. The acknowledgment number is set to `received sequence number A + 1`. And the sequence number that the server chooses for the packet is another random number B.
3. The Client sends an ACK back to the server. The sequence number is set to the recieved acknowledgment value `A + 1`. And the acknowledgment number is set to `received sequence number B + 1`.

Steps 1 and 2 establish and acknowledge the sequence number for one direction. Steps 2 and 3 establish and acknowledge the sequence number for the other direction. Following the completion of these steps, both the client and server have received acknowledgments and a full-duplex communication is established.


## Check Packet Is Received Using Sequence Number

TCP packet that the client sends contains `sequence number`. The sequence number `X + 1` of ACK packet means server receives packets up to `X` and expects to get `X + 1`. Similarly, the server send SYN packet `Y` to the client. And the client send ACK `Y + 1`packet to the server.

![]

For example, the client sends packet with SYN upto 100 but got ACK 90 from the server. The client should retransmit packet with sequence number 91 ~ 100. It is called as `SACK`(Selective Acknowledgement)


## Reference

- https://en.wikipedia.org/wiki/Transmission_Control_Protocol
