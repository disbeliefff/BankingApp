package api

// import (
// 	"bytes"
// 	"database/sql"
// 	"encoding/json"
// 	"errors"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// 	"time"

// 	mockdb "bankingapp/db/mock"
// 	db "bankingapp/db/sqlc"
// 	"bankingapp/token"

// 	"github.com/gin-gonic/gin"
// 	"github.com/golang/mock/gomock"
// 	"github.com/stretchr/testify/require"
// )

// func TestCreateTransferAPI(t *testing.T) {
// 	user, _ := randomUser(t)
// 	account1 := randomAccount(user.Username)
// 	account2 := randomAccount(user.Username)
// 	account3 := randomAccount("other_user")

// 	account1.Currency = "USD"
// 	account2.Currency = "USD"
// 	account3.Currency = "USD"

// 	amount := int64(10)

// 	testCases := []struct {
// 		name          string
// 		body          gin.H
// 		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
// 		buildStubs    func(store *mockdb.MockStore)
// 		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
// 	}{
// 		{
// 			name: "OK",
// 			body: gin.H{
// 				"from_account_id": account1.ID,
// 				"to_account_id":   account2.ID,
// 				"amount":          amount,
// 				"currency":        "USD",
// 			},
// 			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
// 				addAuth(t, request, tokenMaker, authTypeBearer, user.Username, time.Minute)
// 			},
// 			buildStubs: func(store *mockdb.MockStore) {
// 				store.EXPECT().
// 					GetAccount(gomock.Any(), gomock.Eq(account1.ID)).
// 					Times(1).
// 					Return(account1, nil)
// 				store.EXPECT().
// 					GetAccount(gomock.Any(), gomock.Eq(account2.ID)).
// 					Times(1).
// 					Return(account2, nil)

// 				arg := db.TranserTxParams{
// 					FromAccountId: account1.ID,
// 					ToAccountId:   account2.ID,
// 					Amount:        amount,
// 				}

// 				store.EXPECT().
// 					TransferTx(gomock.Any(), gomock.Eq(arg)).
// 					Times(1)
// 			},
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusOK, recorder.Code)
// 			},
// 		},
// 		{
// 			name: "FromAccountNotFound",
// 			body: gin.H{
// 				"from_account_id": account1.ID,
// 				"to_account_id":   account2.ID,
// 				"amount":          amount,
// 				"currency":        "USD",
// 			},
// 			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
// 				addAuth(t, request, tokenMaker, authTypeBearer, user.Username, time.Minute)
// 			},
// 			buildStubs: func(store *mockdb.MockStore) {
// 				store.EXPECT().
// 					GetAccount(gomock.Any(), gomock.Eq(account1.ID)).
// 					Times(1).
// 					Return(db.Account{}, sql.ErrNoRows)
// 				store.EXPECT().
// 					GetAccount(gomock.Any(), gomock.Eq(account2.ID)).
// 					Times(0)
// 				store.EXPECT().
// 					TransferTx(gomock.Any(), gomock.Any()).
// 					Times(0)
// 			},
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusNotFound, recorder.Code)
// 			},
// 		},
// 		{
// 			name: "ToAccountNotFound",
// 			body: gin.H{
// 				"from_account_id": account1.ID,
// 				"to_account_id":   account2.ID,
// 				"amount":          amount,
// 				"currency":        "USD",
// 			},
// 			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
// 				addAuth(t, request, tokenMaker, authTypeBearer, user.Username, time.Minute)
// 			},
// 			buildStubs: func(store *mockdb.MockStore) {
// 				store.EXPECT().
// 					GetAccount(gomock.Any(), gomock.Eq(account1.ID)).
// 					Times(1).
// 					Return(account1, nil)
// 				store.EXPECT().
// 					GetAccount(gomock.Any(), gomock.Eq(account2.ID)).
// 					Times(1).
// 					Return(db.Account{}, sql.ErrNoRows)
// 				store.EXPECT().
// 					TransferTx(gomock.Any(), gomock.Any()).
// 					Times(0)
// 			},
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusNotFound, recorder.Code)
// 			},
// 		},
// 		{
// 			name: "CurrencyMismatch",
// 			body: gin.H{
// 				"from_account_id": account1.ID,
// 				"to_account_id":   account2.ID,
// 				"amount":          amount,
// 				"currency":        "EUR",
// 			},
// 			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
// 				addAuth(t, request, tokenMaker, authTypeBearer, user.Username, time.Minute)
// 			},
// 			buildStubs: func(store *mockdb.MockStore) {
// 				store.EXPECT().
// 					GetAccount(gomock.Any(), gomock.Eq(account1.ID)).
// 					Times(1).
// 					Return(account1, nil)
// 				store.EXPECT().
// 					GetAccount(gomock.Any(), gomock.Eq(account2.ID)).
// 					Times(1).
// 					Return(account2, nil)
// 				store.EXPECT().
// 					TransferTx(gomock.Any(), gomock.Any()).
// 					Times(0)
// 			},
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusBadRequest, recorder.Code)
// 			},
// 		},
// 		{
// 			name: "InternalError",
// 			body: gin.H{
// 				"from_account_id": account1.ID,
// 				"to_account_id":   account2.ID,
// 				"amount":          amount,
// 				"currency":        "USD",
// 			},
// 			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
// 				addAuth(t, request, tokenMaker, authTypeBearer, user.Username, time.Minute)
// 			},
// 			buildStubs: func(store *mockdb.MockStore) {
// 				store.EXPECT().
// 					GetAccount(gomock.Any(), gomock.Eq(account1.ID)).
// 					Times(1).
// 					Return(account1, nil)
// 				store.EXPECT().
// 					GetAccount(gomock.Any(), gomock.Eq(account2.ID)).
// 					Times(1).
// 					Return(account2, nil)
// 				arg := db.TranserTxParams{
// 					FromAccountId: account1.ID,
// 					ToAccountId:   account2.ID,
// 					Amount:        amount,
// 				}
// 				store.EXPECT().
// 					TransferTx(gomock.Any(), gomock.Eq(arg)).
// 					Times(1).
// 					Return(db.TransferTxResult{}, errors.New("internal error"))
// 			},
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusInternalServerError, recorder.Code)
// 			},
// 		},
// 	}

// 	for _, tc := range testCases {
// 		tc := tc // capture range variable
// 		t.Run(tc.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()

// 			store := mockdb.NewMockStore(ctrl)
// 			tc.buildStubs(store)

// 			server := newTestServer(t, store)
// 			recorder := httptest.NewRecorder()

// 			url := "/transfers"
// 			data, err := json.Marshal(tc.body)
// 			require.NoError(t, err)

// 			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
// 			require.NoError(t, err)

// 			tc.setupAuth(t, request, server.tokenMaker)
// 			server.router.ServeHTTP(recorder, request)
// 			tc.checkResponse(t, recorder)
// 		})
// 	}
// }
