module github.com/majidkarimizadeh/sdk-modules/v1

go 1.22.9

require (
    github.com/majidkarimizadeh/sdk-modules/abuse master
    github.com/majidkarimizadeh/sdk-modules/aggregationpack master
    github.com/majidkarimizadeh/sdk-modules/dedicatedserver master
    github.com/majidkarimizadeh/sdk-modules/invoice master
    github.com/majidkarimizadeh/sdk-modules/dns master
    github.com/majidkarimizadeh/sdk-modules/publiccloud master
)

replace github.com/majidkarimizadeh/sdk-modules/abuse => ./abuse
replace github.com/majidkarimizadeh/sdk-modules/aggregationpack => ./aggregationpack
replace github.com/majidkarimizadeh/sdk-modules/dedicatedserver => ./dedicatedserver
replace github.com/majidkarimizadeh/sdk-modules/invoice => ./invoice
replace github.com/majidkarimizadeh/sdk-modules/dns => ./dns
replace github.com/majidkarimizadeh/sdk-modules/publiccloud => ./publiccloud
