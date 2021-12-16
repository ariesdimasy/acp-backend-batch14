# System Design

## Flow Chart

## Use Case

## ERD - Entitry Relationship Diagram

## High Level Architecture

### Load Balancer

adalah implementasi untuk menampung request dan membagikannya secara rata terhadap server

# Key characteristics of Distributed System

## Scaling ( Scalability )

bagaimana cara agar sistem bertumbuh

### A.Horizontal Scaling

menambah jumlah server dengan kemampuan yang sama. penambahan

### B. Vertical Scaling

System yang meningkay kemampuannya. penambahan resources

penentuan dari penggunaan scale tersebut bisa ditentukan dari budget. jika budget sedikit biasanya lebih memilih horizontal scale

## Realibility

## Availability

## Eficiency

## Messsage Broker

### Pub Sub

ada Publisher langsung kirim kerjaan ke subscriber.
contoh Pub Sub adalah :

- google pub sub
- rabbit MQ

### Job / Work Queue

ada producer --> job queue --> Worker

## Load Balancer

## Monolothic

System yang di buat dalam satu kesatuan

## Microservice

- Front End dan Backend Terpisah
- memisahkan 1 backend dengan Backend yang lain
- microservice membutuhkan lebih banyak orang dalam pengembangan

link ke microservice
[Programmer Zaman Now - Microservice ](https://www.youtube.com/watch?v=PhCrso3J11k&list=PL-CtdCApEFH-MtoBwQ0F3xNG21yjt5Kvs)
