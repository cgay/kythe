java_test(
    name = "kzip_reader_test",
    test_class = "com.google.devtools.kythe.platform.kzip.KZipReaderTest",
    srcs = ["KZipReaderTest.java"],
    data = [
            "//kythe/cxx/common/testdata:empty.kzip",
            "//kythe/cxx/common/testdata:garbage_unit.kzip",
            "//kythe/cxx/common/testdata:malformed.kzip",
            "//kythe/cxx/common/testdata:stringset.kzip",
            "//kythe/cxx/common/testdata:stringset_with_empty_file.kzip",
    ],
    deps = [
        ":test_util",
        "//kythe/java/com/google/devtools/kythe/platform/kzip:kzip",
        "//kythe/java/com/google/devtools/kythe/util:json",
        "//kythe/proto:analysis_java_proto",
        "//kythe/proto:go_java_proto",
        "//third_party/truth",
        "@com_google_guava_guava//jar",
        "@com_google_code_gson_gson//jar",
        "@com_google_protobuf//:protobuf_java_util",
        "@junit_junit//jar",
    ],
)

java_test(
    name = "kzip_writer_test",
    test_class = "com.google.devtools.kythe.platform.kzip.KZipWriterTest",
    srcs = ["KZipWriterTest.java"],
    data = ["//kythe/cxx/common/testdata:stringset.kzip"],
    deps = [
        ":test_util",
        "//kythe/java/com/google/devtools/kythe/platform/kzip:kzip",
        "//kythe/java/com/google/devtools/kythe/util:json",
        "//kythe/proto:analysis_java_proto",
        "//kythe/proto:go_java_proto",
        "//third_party/truth",
        "@com_google_protobuf//:protobuf_java_util",
        "@junit_junit//jar",
    ],
)

java_library(
    name = "test_util",
    srcs = ["TestDataUtil.java"],
)
