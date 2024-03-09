use warp::Filter;

#[tokio::main]
async fn main() {
    let index_html = warp::path::end()
        .and(warp::fs::file("index.html"))
        .with(warp::reply::with::header("content-type", "text/html"));

    let static_files = warp::path("static").and(warp::fs::dir("./static"));

    let routes = static_files.or(index_html);

    warp::serve(routes).run(([127, 0, 0, 1], 3030)).await;
}