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