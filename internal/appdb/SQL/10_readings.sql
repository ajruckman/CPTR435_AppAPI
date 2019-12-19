create table readings
(
    id      serial primary key,
    time    timestamp default now(),
    reading float4
);