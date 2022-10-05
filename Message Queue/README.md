# Message Queue
## Queue
| Tightly Coupled | Loosely Coupled |
| --------------- | --------------- |
| ![tight](https://i.imgur.com/Vfpk4bC.png) | ![loose](https://i.imgur.com/R9S1rSu.png ) |  

 Queue is a great data structure for decoupling components and layers, meaning that they can work independently without knowing each other. Benefits such as increased testability, better reliability, asynchronous tasks will follow. By monitoring queue with its size and age, consumers can also auto-scale and manage their processes' state.
 - **Asynchronous**: Producers and consumers only depend on queues, not directly each other
 - **Testability**: Dependencies can sink down into interfaces and data models, allowing easy mocking
 - **Auto Scaling**: Based on queue's size and age, consumers can manage their jobs
 - **Reliability**: As long as queue is alive, system will work correctly even if one side is down

## AMQP
![AMQP](https://media-exp1.licdn.com/dms/image/C4E12AQGMbD9fqiwfLw/article-cover_image-shrink_720_1280/0/1521942243982?e=1670457600&v=beta&t=9JJLP5jQmfhPMM9_rMyylj4ZrAta6XFV_9NlqAgr0Ck)  
