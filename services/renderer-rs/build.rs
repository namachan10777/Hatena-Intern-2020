use protobuf_codegen_pure::Customize;

fn main() {
    protobuf_codegen_pure::Codegen::new()
        .customize(Customize {
            ..Default::default()
        })
        .out_dir("src/protos")
        .input("../../pb/renderer.proto")
        .include("../../pb")
        .run()
        .expect("protoc");
}
