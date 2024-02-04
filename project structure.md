clean-go/
├──  cmd/
           └── main.go
├──  api/
      └── api-spec.json(postman api collection)
      
├──  internal/
      ├──  controller/
        └── user_handlers.go(layer controller handle http request)
 ├──  service/
        └── user_service.go(layer service handle business logic)
 ├──  repositories/
        └── user-repository.go(layer repositories handle http interaksi ke db)
 ├──  entities/
        └── user.go(decklarasi struct tabel)
 ├──  middleware/
        └── middleware.go(kode middleware disini)
      ├──  routes/
        └── routes.go(inisialisasi route untuk server dan run server disini)
 ├──  migration/
        └── User.go(inisialisasi tabel user untuk migrasi ke db)
        └── migrate.go(setup migrate file)
        └── token.go(inisialisasi tabel token untuk migrasi ke db)
├──  config/
        └── db.go(inisialisasi db disini)
├──  test/
        └── user_test.go(unit testing disini)
└──  go.mod
└──  makefile
└──  .env
├──  vendor/