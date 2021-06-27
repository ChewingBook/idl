/* eslint-disable */
import { util, configure, Reader, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Timestamp } from "../google/protobuf/timestamp";

export const protobufPackage = "";

export interface ListArticlesRequest {}

export interface ListArticlesResponse {
  articles: Article[];
}

/** 유저가 작성한 글 */
export interface Article {
  /** 고유식별자 */
  id: number;
  /** 글의 제목 */
  title: string;
  /** 글의 내용 */
  content: string;
  /** 글의 요약 */
  summary: string;
  /** 글의 저자 */
  authorName: string;
  /** 최초로 만들어진 날짜 */
  createdTime: Date | undefined;
  /** 마지막으로 수정된 날짜 */
  updatedTime: Date | undefined;
  /** 글을 볼 수 있는 웹페이지 주소 */
  url: string;
}

const baseListArticlesRequest: object = {};

export const ListArticlesRequest = {
  encode(_: ListArticlesRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ListArticlesRequest {
    const reader = input instanceof Reader ? input : new Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseListArticlesRequest } as ListArticlesRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): ListArticlesRequest {
    const message = { ...baseListArticlesRequest } as ListArticlesRequest;
    return message;
  },

  toJSON(_: ListArticlesRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<ListArticlesRequest>): ListArticlesRequest {
    const message = { ...baseListArticlesRequest } as ListArticlesRequest;
    return message;
  },
};

const baseListArticlesResponse: object = {};

export const ListArticlesResponse = {
  encode(
    message: ListArticlesResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.articles) {
      Article.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ListArticlesResponse {
    const reader = input instanceof Reader ? input : new Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseListArticlesResponse } as ListArticlesResponse;
    message.articles = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.articles.push(Article.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ListArticlesResponse {
    const message = { ...baseListArticlesResponse } as ListArticlesResponse;
    message.articles = [];
    if (object.articles !== undefined && object.articles !== null) {
      for (const e of object.articles) {
        message.articles.push(Article.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: ListArticlesResponse): unknown {
    const obj: any = {};
    if (message.articles) {
      obj.articles = message.articles.map((e) =>
        e ? Article.toJSON(e) : undefined
      );
    } else {
      obj.articles = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<ListArticlesResponse>): ListArticlesResponse {
    const message = { ...baseListArticlesResponse } as ListArticlesResponse;
    message.articles = [];
    if (object.articles !== undefined && object.articles !== null) {
      for (const e of object.articles) {
        message.articles.push(Article.fromPartial(e));
      }
    }
    return message;
  },
};

const baseArticle: object = {
  id: 0,
  title: "",
  content: "",
  summary: "",
  authorName: "",
  url: "",
};

export const Article = {
  encode(message: Article, writer: Writer = Writer.create()): Writer {
    if (message.id !== 0) {
      writer.uint32(8).int64(message.id);
    }
    if (message.title !== "") {
      writer.uint32(18).string(message.title);
    }
    if (message.content !== "") {
      writer.uint32(26).string(message.content);
    }
    if (message.summary !== "") {
      writer.uint32(34).string(message.summary);
    }
    if (message.authorName !== "") {
      writer.uint32(42).string(message.authorName);
    }
    if (message.createdTime !== undefined) {
      Timestamp.encode(
        toTimestamp(message.createdTime),
        writer.uint32(50).fork()
      ).ldelim();
    }
    if (message.updatedTime !== undefined) {
      Timestamp.encode(
        toTimestamp(message.updatedTime),
        writer.uint32(58).fork()
      ).ldelim();
    }
    if (message.url !== "") {
      writer.uint32(66).string(message.url);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Article {
    const reader = input instanceof Reader ? input : new Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseArticle } as Article;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.int64() as Long);
          break;
        case 2:
          message.title = reader.string();
          break;
        case 3:
          message.content = reader.string();
          break;
        case 4:
          message.summary = reader.string();
          break;
        case 5:
          message.authorName = reader.string();
          break;
        case 6:
          message.createdTime = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        case 7:
          message.updatedTime = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        case 8:
          message.url = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Article {
    const message = { ...baseArticle } as Article;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    if (object.title !== undefined && object.title !== null) {
      message.title = String(object.title);
    } else {
      message.title = "";
    }
    if (object.content !== undefined && object.content !== null) {
      message.content = String(object.content);
    } else {
      message.content = "";
    }
    if (object.summary !== undefined && object.summary !== null) {
      message.summary = String(object.summary);
    } else {
      message.summary = "";
    }
    if (object.authorName !== undefined && object.authorName !== null) {
      message.authorName = String(object.authorName);
    } else {
      message.authorName = "";
    }
    if (object.createdTime !== undefined && object.createdTime !== null) {
      message.createdTime = fromJsonTimestamp(object.createdTime);
    } else {
      message.createdTime = undefined;
    }
    if (object.updatedTime !== undefined && object.updatedTime !== null) {
      message.updatedTime = fromJsonTimestamp(object.updatedTime);
    } else {
      message.updatedTime = undefined;
    }
    if (object.url !== undefined && object.url !== null) {
      message.url = String(object.url);
    } else {
      message.url = "";
    }
    return message;
  },

  toJSON(message: Article): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.title !== undefined && (obj.title = message.title);
    message.content !== undefined && (obj.content = message.content);
    message.summary !== undefined && (obj.summary = message.summary);
    message.authorName !== undefined && (obj.authorName = message.authorName);
    message.createdTime !== undefined &&
      (obj.createdTime = message.createdTime.toISOString());
    message.updatedTime !== undefined &&
      (obj.updatedTime = message.updatedTime.toISOString());
    message.url !== undefined && (obj.url = message.url);
    return obj;
  },

  fromPartial(object: DeepPartial<Article>): Article {
    const message = { ...baseArticle } as Article;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    if (object.title !== undefined && object.title !== null) {
      message.title = object.title;
    } else {
      message.title = "";
    }
    if (object.content !== undefined && object.content !== null) {
      message.content = object.content;
    } else {
      message.content = "";
    }
    if (object.summary !== undefined && object.summary !== null) {
      message.summary = object.summary;
    } else {
      message.summary = "";
    }
    if (object.authorName !== undefined && object.authorName !== null) {
      message.authorName = object.authorName;
    } else {
      message.authorName = "";
    }
    if (object.createdTime !== undefined && object.createdTime !== null) {
      message.createdTime = object.createdTime;
    } else {
      message.createdTime = undefined;
    }
    if (object.updatedTime !== undefined && object.updatedTime !== null) {
      message.updatedTime = object.updatedTime;
    } else {
      message.updatedTime = undefined;
    }
    if (object.url !== undefined && object.url !== null) {
      message.url = object.url;
    } else {
      message.url = "";
    }
    return message;
  },
};

export interface Articles {
  ListArticles(request: ListArticlesRequest): Promise<ListArticlesResponse>;
}

export class ArticlesClientImpl implements Articles {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.ListArticles = this.ListArticles.bind(this);
  }
  ListArticles(request: ListArticlesRequest): Promise<ListArticlesResponse> {
    const data = ListArticlesRequest.encode(request).finish();
    const promise = this.rpc.request("Articles", "ListArticles", data);
    return promise.then((data) =>
      ListArticlesResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

type Builtin =
  | Date
  | Function
  | Uint8Array
  | string
  | number
  | boolean
  | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function toTimestamp(date: Date): Timestamp {
  const seconds = date.getTime() / 1_000;
  const nanos = (date.getTime() % 1_000) * 1_000_000;
  return { seconds, nanos };
}

function fromTimestamp(t: Timestamp): Date {
  let millis = t.seconds * 1_000;
  millis += t.nanos / 1_000_000;
  return new Date(millis);
}

function fromJsonTimestamp(o: any): Date {
  if (o instanceof Date) {
    return o;
  } else if (typeof o === "string") {
    return new Date(o);
  } else {
    return fromTimestamp(Timestamp.fromJSON(o));
  }
}

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

// If you get a compile-error about 'Constructor<Long> and ... have no overlap',
// add '--ts_proto_opt=esModuleInterop=true' as a flag when calling 'protoc'.
if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
