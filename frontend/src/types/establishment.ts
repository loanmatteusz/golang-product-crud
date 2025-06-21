export type EstablishmentResponse = {
  ID: string;
  Number: string;
  Name: string;
  LegalName: string;
  AddressID: string;
  Address: {
    ID: string;
    Street: string;
    Number: string;
    City: string;
    State: string;
    ZipCode: string;
    CreatedAt: Date;
    UpdatedAt: Date;
  },
  CreatedAt: Date;
  UpdatedAt: Date;
}

export type EstablishmentPreview = {
  id: string;
  name: string;
  legal_name: string;
  number: string;
  address: {
    street: string;
    number: string;
    city: string;
    state: string;
    zip_code: string;
  }
}
