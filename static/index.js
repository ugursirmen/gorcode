(function () {
  $("#table").DataTable({
    ajax: {
      url: "http://localhost:8080/products",
      dataSrc: "",
    },
    columns: [{ data: "code" }, { data: "name" }, { data: "barcode" }],
  });
})();

function addProduct() {}
