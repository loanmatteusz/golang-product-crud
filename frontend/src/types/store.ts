import type { EstablishmentResponse } from "./establishment";

export type StoreResponse = {
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
  EstablishmentID: string;
  Establishment: EstablishmentResponse;
  CreatedAt: Date;
  UpdatedAt: Date;
}

export type StorePreview = {
  id: string;
  name: string;
  legal_name: string;
  number: string;
  establishment_id: string; // Establishment Name
  address: {
    street: string;
    number: string;
    city: string;
    state: string;
    zip_code: string;
  }
}
