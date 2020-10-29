(function () {
  var table = $("#table").DataTable({
    ajax: {
      url: "http://localhost:8080/products",
      dataSrc: "",
    },
    columns: [
      {
        data: "ID",
      },
      {
        data: "barcode",
      },
      { data: "code" },
      { data: "name" },
      {
        data: "CreatedAt",
        render: function (data) {
          return moment(data).format("DD.MM.YYYY HH:mm A");
        }
      },
      {
        data: "UpdatedAt",
        render: function (data) {
          return moment(data).format("DD.MM.YYYY HH:mm A");
        }
      },
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
    $("#selectedId").val(data.ID);

    $("#productModal").modal("show");
  });

  $("#table tbody").on("click", "button.delete-button", function () {
    var data = table.row($(this).parents("tr")).data();

    $("#selectedId").val(data.ID);

    $("#codeLabel").text(data.code);
    $("#nameLabel").text(data.name);
    $("#barcodeLabel").text(data.barcode);

    $("#deleteConfirmationModal").modal("show");
  });

  $("#productModal").on("hidden.bs.modal", function (e) {
    $("#code").val("");
    $("#name").val("");
    $("#selectedId").val("");
  });

  $("#deleteConfirmationModal").on("hidden.bs.modal", function (e) {
    $("#selectedId").val("");
  });

  $("#btnSave").on("click", function () {
    var id = $("#selectedId").val();
    if (!id) {
      addProduct();
    } else {
      updateProduct(id);
    }
  });

  $("#btnDelete").on("click", function () {
    var id = $("#selectedId").val();
    deleteProduct(id);
  });
})();

function addProduct() {
  $.ajax({
    url: "http://localhost:8080/products",
    type: "POST",
    data: JSON.stringify({
      code: $("#code").val(),
      name: $("#name").val(),
    }),
  })
    .done(function (result) {
      console.log(result);
      $("#productModal").modal("hide");
      $("#table").DataTable().ajax.reload();
    })
    .fail(function (xhr, result, status) {
      console.log("Error:", result);
    });
}

function updateProduct(id) {
  $.ajax({
    url: "http://localhost:8080/products/" + id,
    type: "PUT",
    data: JSON.stringify({
      code: $("#code").val(),
      name: $("#name").val(),
    }),
  })
    .done(function (result) {
      console.log(result);
      $("#productModal").modal("hide");
      $("#table").DataTable().ajax.reload();
    })
    .fail(function (xhr, result, status) {
      console.log("Error:", result);
    });
}

function deleteProduct(id) {
  $.ajax({
    url: "http://localhost:8080/products/" + id,
    type: "DELETE",
  })
    .done(function (result) {
      console.log(result);
      $("#deleteConfirmationModal").modal("hide");
      $("#table").DataTable().ajax.reload();
    })
    .fail(function (xhr, result, status) {
      console.log("Error:", result);
    });
}
