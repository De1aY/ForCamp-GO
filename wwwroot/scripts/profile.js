let RequestUserID = null;
try {
    RequestUserID = new URL(window.location).searchParams.get("id");
} catch (e) {
    RequestUserID = GetURLParameter("id");
}