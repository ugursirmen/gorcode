(function () {
  var table = $("#table").DataTable({
    ajax: {
      url: "http://localhost:8080/products",
      dataSrc: "",
    },
    columns: [
      {
        data: "barcode",
      },
      { data: "code" },
      { data: "name" },
      { data: "modifiedAt" },
      {
        targets: -1,
        data: null,
        defaultContent:
          "<button type='button' class='btn btn-sm btn-link edit-button'><i class='fa fa-edit'></i></button> | <button type='button' class='btn btn-sm btn-link delete-button'><i class='fa fa-trash'></i></button>",
      },
    ],
  });

  $("#table tbody").on("click", "button.edit-button", function () {
    var data = table.row($(this).parents("tr")).data();

    $("#code").val(data.code);
    $("#name").val(data.name);
    $("#selectedProductBarcode").val(data.barcode);

    $("#productModal").modal("show");
  });

  $("#table tbody").on("click", "button.delete-button", function () {
    var data = table.row($(this).parents("tr")).data();

    $("#selectedProductBarcode").val(data.barcode);

    $("#codeLabel").text(data.code);
    $("#nameLabel").text(data.name);
    $("#barcodeLabel").text(data.barcode);

    $("#deleteConfirmationModal").modal("show");
  });

  $("#productModal").on("hidden.bs.modal", function (e) {
    $("#code").val("");
    $("#name").val("");
    $("#selectedProductBarcode").val("");
  });

  $("#deleteConfirmationModal").on("hidden.bs.modal", function (e) {
    $("#selectedProductBarcode").val("");
  });

  $("#btnSave").on("click", function () {
    var barcode = $("#selectedProductBarcode").val();
    if (!barcode) {
      addProduct();
    } else {
      updateProduct(barcode);
    }
  });

  $("#btnDelete").on("click", function () {
    var barcode = $("#selectedProductBarcode").val();
    deleteProduct(barcode);
  });

})();

function addProduct() {
  $.ajax({
    url: "http://localhost:8080/products",
    type: "POST",
    data: {
      code: $("#code").val(),
      name: $("#name").val(),
    },
    success: function () {
      $("#productModal").modal("hide");
    },
  });
}

function updateProduct(barcode) {
  $.ajax({
    url: "http://localhost:8080/products/" + barcode,
    type: "PUT",
    data: {
      code: $("#code").val(),
      name: $("#name").val(),
    },
    success: function () {
      $("#productModal").modal("hide");
    },
  });
}

function deleteProduct(barcode) {
  $.ajax({
    url: "http://localhost:8080/products/" + barcode,
    type: "DELETE",
    success: function () {
      $("#deleteConfirmationModal").modal("hide");
    },
  });
}
