# Introduction

This code has problems and does not aims our SLA.

Let's find all the issues and make it great.

# Overview

Program is diveded on 2 parts: data generator and loader. Data generator will generate provided amount of randomly
generate events with following probability:

* type 1 - 15%
* type 2 - 20%
* type 3 - 20&
* type 5 - 45%

All the data must be randomly generated, there are no preferences of generation algorithm.

Events must bo stored on local file system for future loading.

Loading will be done by separate process - loader. That process should insert this data to database.

Our goal is to generate 8M of events and load them in database as soon as it is possible. Hardware - developer machine.

# Table Structure

```sql

create table event
(
    event_source     text not null, -- unique identifier of client
    event_ref        text not null, -- unique identifier of event
    event_type       integer not null,
    event_date       timestamp not null,
    calling_number   BIGINT not null,
    called_number    BIGINT not null,
    location         text not null,
    duration_seconds BIGINT not null,
    attr_1           text,
    attr_2           text,
    attr_3           text,
    attr_4           text,
    attr_5           text,
    attr_6           text,
    attr_7           text,
    attr_8           text,
    PRIMARY KEY (event_source, event_ref)
);


create index event_called_number_index
    on event (called_number);

create index event_calling_number_index
    on event (calling_number);

create index event_event_date_index
    on event (event_date);

create
unique index event_event_ref_uindex
    on event (event_ref);

create index event_event_type_index
    on event (event_type);

create index event_location_index
    on event (location);
```
