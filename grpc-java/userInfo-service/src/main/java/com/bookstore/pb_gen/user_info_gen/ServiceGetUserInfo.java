// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: service_get_user_info.proto

package com.bookstore.pb_gen.user_info_gen;

public final class ServiceGetUserInfo {
  private ServiceGetUserInfo() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    String[] descriptorData = {
      "\n\033service_get_user_info.proto\022\"com.books" +
      "tore.pb_gen.user_info_gen\032\026pb_get_user_i" +
      "nfo.proto2\224\001\n\017UserInfoService\022\200\001\n\013getUse" +
      "rInfo\0226.com.bookstore.pb_gen.user_info_g" +
      "en.getUserInfoRequest\0327.com.bookstore.pb" +
      "_gen.user_info_gen.getUserInfoResponse\"\000" +
      "B\014Z\n/user_infob\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          PbGetUserInfo.getDescriptor(),
        });
    PbGetUserInfo.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
