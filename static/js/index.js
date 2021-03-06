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
        data: "image",
        render: function (image) {
          if (image) {
            return (
              "<img src='data:image/png;base64," +
              image +
              "' class='img-fluid img-thumbnail'/>"
            );
          }

          return "<img src='http://localhost:8080/default.jpg' class='img-fluid img-thumbnail'/>";
        },
        width: "10%",
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
        },
      },
      {
        data: "UpdatedAt",
        render: function (data) {
          return moment(data).format("DD.MM.YYYY HH:mm A");
        },
      },
      {
        targets: -1,
        data: null,
        defaultContent:
          "<button type='button' class='btn btn-sm btn-link edit-button'><i class='fa fa-edit'></i></button> | <button type='button' class='btn btn-sm btn-link delete-button'><i class='fa fa-trash'></i></button> | <button type='button' class='btn btn-sm btn-link barcode-button'><i class='fa fa-barcode'></i></button>",
      },
    ],
    order: [],
    scrollX: true,
    pageLength : 5,
    lengthMenu: [[5, 10, 20, 50, 100], [5, 10, 20, 50, 100]]
  });

  table.columns([0]).visible(false);

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

    $("#deleteConfirmationModal").modal("show");
  });

  $("#table tbody").on("click", "button.barcode-button", function () {
    var data = table.row($(this).parents("tr")).data();

    $.get("http://localhost:8080/barcode/" + data.barcode, function (imgb64) {
      $("#barcodeImage").attr("src", "data:image/png;base64," + imgb64);
    });

    $("#barcodeModal").modal("show");
  });

  $("#productModal").on("hidden.bs.modal", function (e) {
    $("#code").val("");
    $("#name").val("");
    $("#image").val("");
    $("#selectedId").val("");
  });

  $("#deleteConfirmationModal").on("hidden.bs.modal", function (e) {
    $("#selectedId").val("");
  });

  $("#barcodeModal").on("hidden.bs.modal", function (e) {
    $("#barcodeImage").attr("src", "");
  });

  $("#btnSave").on("click", function () {

    if(!$("#productForm").valid()){
      return
    }

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
  $("#btnSave").button('loading');
  $.ajax({
    url: "http://localhost:8080/products",
    type: "POST",
    contentType: false,
    processData: false,
    data: new FormData($('form')[0])
  })
    .done(function (result) {
      console.log(result);
      $("#productModal").modal("hide");
      $("#table").DataTable().ajax.reload();
    })
    .fail(function (xhr, result, status) {
      console.log("Error:", result);
    })
    .always(function(){
      $("#btnSave").button('reset');
    });
}

function updateProduct(id) {
  $.ajax({
    url: "http://localhost:8080/products/" + id,
    type: "PUT",
    contentType: false,
    processData: false,
    data: new FormData($('form')[0])
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

function print() {
  css = new String('<link href="index.css" rel="stylesheet" type="text/css">');
  window.frames["print_frame"].document.body.innerHTML =
    css + document.getElementById("barcodeDiv").innerHTML;
  window.frames["print_frame"].window.focus();
  window.frames["print_frame"].window.print();
}
