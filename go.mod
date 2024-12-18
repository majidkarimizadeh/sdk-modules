module github.com/majidkarimizadeh/sdk-modules/v1

go 1.22.9

replace github.com/majidkarimizadeh/sdk-modules/abuse => ./abuse

replace github.com/majidkarimizadeh/sdk-modules/aggregationpack => ./aggregationpack

replace github.com/majidkarimizadeh/sdk-modules/dedicatedserver => ./dedicatedserver

replace github.com/majidkarimizadeh/sdk-modules/invoice => ./invoice

replace github.com/majidkarimizadeh/sdk-modules/dns => ./dns

replace github.com/majidkarimizadeh/sdk-modules/publiccloud => ./publiccloud

require (
	github.com/majidkarimizadeh/sdk-modules/abuse v0.0.0-00010101000000-000000000000
	github.com/majidkarimizadeh/sdk-modules/aggregationpack v0.0.0-00010101000000-000000000000
	github.com/majidkarimizadeh/sdk-modules/dedicatedserver v0.0.0-00010101000000-000000000000
	github.com/majidkarimizadeh/sdk-modules/dns v0.0.0-00010101000000-000000000000
	github.com/majidkarimizadeh/sdk-modules/invoice v0.0.0-00010101000000-000000000000
	github.com/majidkarimizadeh/sdk-modules/publiccloud v0.0.0-00010101000000-000000000000
)

require gopkg.in/validator.v2 v2.0.1 // indirect
