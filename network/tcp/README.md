# Reliable TCP Datastream

## Transmission Control Protocol

- Ordered data tranfer: the destination host rearranges segments according to sequence number.
- Retransmission of lost packets: any cumulative stream not acknowledged is retransmitted.
- Error-free data transfer: corrupted packtets are treated as lost and are retransmitted.
- Flow control: limits the rate a sender transfer data to gurantee reliable delivery. The receiver continually hints the sender on how much data can be received. When the receiving host's buffer fills, the next acknowlegment suspends the transfer and allows the data in the buffer to be processed.
- Congestion control: lost packets trigger a reduction in data delivery rate.


## TCP Session with 3-Way Handshake

<img width="285" alt="image" src="https://user-images.githubusercontent.com/27891090/200204963-debc965f-2429-4c81-a6fc-68e0a3cfa4f3.png">


1. Active open: The client Sends `SYN` to the server. The client sets the segment's sequence number to a random value A.
2. The server replies with a `SYN-ACK`. The acknowledgment number is set to `received sequence number A + 1`. And the sequence number that the server chooses for the packet is another random number B.
3. The Client sends an ACK back to the server. The sequence number is set to the recieved acknowledgment value `A + 1`. And the acknowledgment number is set to `received sequence number B + 1`.

Steps 1 and 2 establish and acknowledge the sequence number for one direction. Steps 2 and 3 establish and acknowledge the sequence number for the other direction. Following the completion of these steps, both the client and server have received acknowledgments and a full-duplex communication is established.


## Check Packet Is Received Using Sequence Number

TCP packet that the client sends contains `sequence number`. The sequence number `X + 1` of ACK packet means server receives packets up to `X` and expects to get `X + 1`. Similarly, the server send SYN packet `Y` to the client. And the client send ACK `Y + 1`packet to the server.

<img width="269" alt="image" src="https://user-images.githubusercontent.com/27891090/200204976-39c62ebb-5ec5-48b2-aab8-a4b02dff564b.png">

For example, the client sends packet with SYN upto 100 but got ACK 90 from the server. The client should retransmit packet with sequence number 91 ~ 100. It is called as `SACK`(Selective Acknowledgement)


## Receive Buffer And Slide Window Size

Because TCP use only one ACK packet to check more than one packets are received, the receiver should tell the sender how much receive buffer size can be available before the sender receives the ACK packet.
Receive Buffer is memory space to use for receiving data for network. ACK packet contains window size. It means the bytes size that the sender can transmit without checking that the receiver receives packets.

<img width="479" alt="image" src="https://user-images.githubusercontent.com/27891090/200204993-9f589d02-7443-48df-acca-eb0d23d1ca8d.png">

## Graceful Termination in TCP Session

<img width="378" alt="image" src="https://user-images.githubusercontent.com/27891090/200205005-53bad58c-6ba7-4064-9a54-5a5506ca5b46.png">


## Not Graceful Termination in TCP session

RST packet

## Reference

- https://en.wikipedia.org/wiki/Transmission_Control_Protocol
