import * as jspb from 'google-protobuf'



export class User extends jspb.Message {
  getName(): string;
  setName(value: string): User;

  getGender(): string;
  setGender(value: string): User;

  getStatus(): string;
  setStatus(value: string): User;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): User.AsObject;
  static toObject(includeInstance: boolean, msg: User): User.AsObject;
  static serializeBinaryToWriter(message: User, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): User;
  static deserializeBinaryFromReader(message: User, reader: jspb.BinaryReader): User;
}

export namespace User {
  export type AsObject = {
    name: string,
    gender: string,
    status: string,
  }
}

export class AddUserRequest extends jspb.Message {
  getUser(): User | undefined;
  setUser(value?: User): AddUserRequest;
  hasUser(): boolean;
  clearUser(): AddUserRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddUserRequest.AsObject;
  static toObject(includeInstance: boolean, msg: AddUserRequest): AddUserRequest.AsObject;
  static serializeBinaryToWriter(message: AddUserRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddUserRequest;
  static deserializeBinaryFromReader(message: AddUserRequest, reader: jspb.BinaryReader): AddUserRequest;
}

export namespace AddUserRequest {
  export type AsObject = {
    user?: User.AsObject,
  }
}

export class AddUserResponse extends jspb.Message {
  getStatus(): string;
  setStatus(value: string): AddUserResponse;

  getUser(): User | undefined;
  setUser(value?: User): AddUserResponse;
  hasUser(): boolean;
  clearUser(): AddUserResponse;

  getError(): string;
  setError(value: string): AddUserResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddUserResponse.AsObject;
  static toObject(includeInstance: boolean, msg: AddUserResponse): AddUserResponse.AsObject;
  static serializeBinaryToWriter(message: AddUserResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddUserResponse;
  static deserializeBinaryFromReader(message: AddUserResponse, reader: jspb.BinaryReader): AddUserResponse;
}

export namespace AddUserResponse {
  export type AsObject = {
    status: string,
    user?: User.AsObject,
    error: string,
  }
}

