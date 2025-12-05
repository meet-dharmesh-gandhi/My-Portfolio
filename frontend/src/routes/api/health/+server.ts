import { BACKEND_URL } from "$env/static/private";
import { json } from "@sveltejs/kit";

export async function GET() {
    let response = await fetch(`${BACKEND_URL}/utils/health`).then(data => data.json()).then((data) => {
        return {
            "status": "success",
            data,
        }
    }).catch(e => {
        console.log("Error while fetching health...");
        console.log(e);
        return {
            "status": "error",
            data: e
        };
    });

    return json(response);
}
