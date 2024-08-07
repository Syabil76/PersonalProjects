Table Garage {
  id bigserial [pk]
  car_id bigint [ref: > C.id]
  brand_id bigint [ref: > C.Make]
  owner_id bigint [ref: > O.id]
}

Table Cars as C {
  id bigserial [pk]
  Make varchar [ref: > B.Name]
  Model varchar
  Horsepower bigint
  Bought_at timestamptz [default: 'now()']
}

Table Brand as B{
  id bugserial [pk]
  Name varchar
  Country varchar

}

Table Owner as O{
  id bigserial [pk]
  FirstName varchar
  LastName varchar
  Age int
  Address varchar
}
