# Message Queue
## Queue
| Tightly Coupled | Loosely Coupled |
| --------------- | --------------- |
| ![tight](https://i.imgur.com/Vfpk4bC.png) | ![loose](https://i.imgur.com/R9S1rSu.png ) |  

 Queue is a great data structure for decoupling components and layers, meaning that they can work independently without knowing each other. Benefits such as increased testability, better reliability, asynchronous tasks will follow. By monitoring queue with its **size** and **age**, consumers can also auto-scale and manage their processes' state.
 - **Asynchronous**: Producers and consumers only depend on queues, not directly each other
 - **Testability**: Dependencies can sink down into interfaces and data models, allowing easy mocking
 - **Auto Scaling**: Based on queue's size and age, consumers can manage their jobs
 - **Reliability**: As long as queue is alive, system will work correctly even if one side is down

## AMQP
![AMQP](https://paolopatierno.files.wordpress.com/2015/07/amqp_frame.png) 
[Advanced Message Queue Protocol](https://www.rabbitmq.com/tutorials/amqp-concepts.html) is an application layer protocol that focuses on process-to-process communication across IP networks. It defines a message structure and transport procedures. Message is split into 3 parts:
- **Header**: AMQP specifications (Key-Value pairs)
- **Properties**: Application specific information (Key-Value pairs)
- **Body**:  Payload
