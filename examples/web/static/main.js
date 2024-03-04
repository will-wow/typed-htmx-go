function main() {
  function handleHtmxErrors() {
    document.body.addEventListener("htmx:beforeSwap", function (evt) {
      if (evt.detail.xhr.status === 404) {
        // alert the user when a 404 occurs (maybe use a nicer mechanism than alert())
        alert("Error: Could Not Find Resource");
      } else if (evt.detail.xhr.status === 422) {
        // allow 422 responses to swap as we are using this as a signal that
        // a form was submitted with bad data and want to rerender with the
        // errors
        //
        // set isError to false to avoid error logging in console
        evt.detail.shouldSwap = true;
        evt.detail.isError = false;
      }
    });
  }

  function highlight() {
    htmx.onLoad(function (el) {
      // highlight any new code blocks
      document.querySelectorAll("pre code:not(.hljs)").forEach((el) => {
        hljs.highlightElement(el);
      });
    });
  }

  highlight();
  handleHtmxErrors();
}

main();
