# Benchmark Golang and Nodejs Parquet Performance

## Getting Started

### Dependencies

* golang - 1.15.6
* nodejs - 17.2.0

### Data


```
{"media_account_uid":"1529888140439425","media_account":"Footer","campaign_id":"23843461555820022","campaign":"(X\/\u5e38\u614b\u671f\u9593\/CPA\/\u896a\u5b50_\u820a\u5ba2_CPA)\u896a\u5b50_\u820a\u5ba2_CPA#[CUE00008182][Facebook][Footer][X][X][X][Banner][CPA]","adset_id":"23843461571790022","adset":"(RT\/\u52a0\u5165\u8cfc\u7269\u8eca\/X\/30days+member30days)cart30days+member30days_all#[X]","ad_id":"23843461710790022","ad":"(X\/\u6a5f\u80fd\/X\/Star\/\u6a5f\u80fd\u896a\/\u751f\u6d3b\u60c5\u5883\/HP\/\u8f2a\u64ad_\u89e3\u6c7a\u5c08\u5bb6)\u8f2a\u64ad_\u9664\u81ed\u7cfb\u5217_RT_190603#[Banner][1080x1080]","ad_type_id":1,"device_id":1,"day":"2020-01-01","action_values_offsite_conversion_fb_pixel_view_content_value":380,"action_values_offsite_conversion_fb_pixel_add_to_cart_value":1240,"action_values_omni_view_content_value":380,"action_values_omni_add_to_cart_value":1240,"actions_offsite_conversion_fb_pixel_view_content_value":1,"actions_offsite_conversion_fb_pixel_add_to_cart_value":3,"actions_link_click_value":1,"actions_landing_page_view_value":1,"actions_omni_view_content_value":1,"actions_omni_add_to_cart_value":3,"actions_page_engagement_value":1,"actions_post_engagement_value":1,"clicks":1,"frequency":4.059072,"impressions":962,"inline_link_clicks":1,"inline_post_engagement":1,"reach":237,"spend":42,"cost":42}
{"media_account_uid":"1529888140439425","media_account":"Footer","campaign_id":"23843461555820022","campaign":"(X\/\u5e38\u614b\u671f\u9593\/CPA\/\u896a\u5b50_\u820a\u5ba2_CPA)\u896a\u5b50_\u820a\u5ba2_CPA#[CUE00008182][Facebook][Footer][X][X][X][Banner][CPA]","adset_id":"23843461571790022","adset":"(RT\/\u52a0\u5165\u8cfc\u7269\u8eca\/X\/30days+member30days)cart30days+member30days_all#[X]","ad_id":"23843461710790022","ad":"(X\/\u6a5f\u80fd\/X\/Star\/\u6a5f\u80fd\u896a\/\u751f\u6d3b\u60c5\u5883\/HP\/\u8f2a\u64ad_\u89e3\u6c7a\u5c08\u5bb6)\u8f2a\u64ad_\u9664\u81ed\u7cfb\u5217_RT_190603#[Banner][1080x1080]","ad_type_id":1,"device_id":2,"day":"2020-01-01","action_values_offsite_conversion_fb_pixel_purchase_value":6475,"action_values_offsite_conversion_fb_pixel_view_content_value":13260,"action_values_offsite_conversion_fb_pixel_initiate_checkout_value":1080,"action_values_offsite_conversion_fb_pixel_add_to_cart_value":7590,"action_values_offsite_conversion_fb_pixel_add_payment_info_value":720,"action_values_omni_view_content_value":13260,"action_values_omni_search_value":0,"action_values_omni_purchase_value":6475,"action_values_lead_value":0,"action_values_omni_initiated_checkout_value":1080,"action_values_omni_complete_registration_value":0,"action_values_omni_add_to_cart_value":7590,"action_values_add_payment_info_value":720,"actions_offsite_conversion_fb_pixel_purchase_value":3,"actions_offsite_conversion_fb_pixel_view_content_value":35,"actions_offsite_conversion_fb_pixel_search_value":1,"actions_offsite_conversion_fb_pixel_lead_value":7,"actions_offsite_conversion_fb_pixel_initiate_checkout_value":4,"actions_offsite_conversion_fb_pixel_complete_registration_value":1,"actions_offsite_conversion_fb_pixel_add_to_cart_value":24,"actions_offsite_conversion_fb_pixel_add_payment_info_value":2,"actions_link_click_value":11,"actions_landing_page_view_value":11,"actions_omni_view_content_value":35,"actions_omni_search_value":1,"actions_omni_purchase_value":3,"actions_lead_value":7,"actions_omni_initiated_checkout_value":4,"actions_omni_complete_registration_value":1,"actions_omni_add_to_cart_value":24,"actions_add_payment_info_value":2,"actions_page_engagement_value":11,"actions_post_engagement_value":11,"clicks":18,"frequency":2.890971,"impressions":1648,"inline_link_clicks":11,"inline_post_engagement":11,"purchase_roas_omni_purchase_value":14.199561,"reach":1184,"spend":459,"website_purchase_roas_offsite_conversion_fb_pixel_purchase_value":14.199561,"cost":459,"conversions":3,"results":3}
{"media_account_uid":"1529888140439425","media_account":"Footer","campaign_id":"23843461555820022","campaign":"(X\/\u5e38\u614b\u671f\u9593\/CPA\/\u896a\u5b50_\u820a\u5ba2_CPA)\u896a\u5b50_\u820a\u5ba2_CPA#[CUE00008182][Facebook][Footer][X][X][X][Banner][CPA]","adset_id":"23843461571790022","adset":"(RT\/\u52a0\u5165\u8cfc\u7269\u8eca\/X\/30days+member30days)cart30days+member30days_all#[X]","ad_id":"23843461710800022","ad":"(X\/\u6a5f\u80fd\/X\/Star\/\u6a5f\u80fd\u896a\/\u5f37\u8abf\u6a5f\u80fd\/HP\/\u55ae\u5716_\u9b54\u93e1_\u8f15\u9b06\u53bb\u5473)\u55ae\u5716_\u9664\u81ed\u7cfb\u5217_RT_190603#[Banner][1080x1080]","ad_type_id":1,"device_id":1,"day":"2020-01-01","action_values_offsite_conversion_fb_pixel_view_content_value":2740,"action_values_offsite_conversion_fb_pixel_purchase_value":1030,"action_values_offsite_conversion_fb_pixel_initiate_checkout_value":1040,"action_values_offsite_conversion_fb_pixel_add_to_cart_value":6240,"action_values_omni_view_content_value":2740,"action_values_omni_purchase_value":1030,"action_values_omni_initiated_checkout_value":1040,"action_values_omni_complete_registration_value":0,"action_values_omni_add_to_cart_value":6240,"actions_offsite_conversion_fb_pixel_view_content_value":8,"actions_offsite_conversion_fb_pixel_purchase_value":1,"actions_offsite_conversion_fb_pixel_initiate_checkout_value":3,"actions_offsite_conversion_fb_pixel_complete_registration_value":1,"actions_offsite_conversion_fb_pixel_add_to_cart_value":19,"actions_link_click_value":2,"actions_landing_page_view_value":2,"actions_omni_view_content_value":8,"actions_omni_purchase_value":1,"actions_omni_initiated_checkout_value":3,"actions_omni_complete_registration_value":1,"actions_omni_add_to_cart_value":19,"actions_page_engagement_value":2,"actions_post_engagement_value":2,"clicks":3,"frequency":2.22467,"impressions":505,"inline_link_clicks":2,"inline_post_engagement":2,"purchase_roas_omni_purchase_value":13.733333,"reach":227,"spend":75,"website_purchase_roas_offsite_conversion_fb_pixel_purchase_value":13.733333,"cost":75,"conversions":1,"results":1}
{"media_account_uid":"1529888140439425","media_account":"Footer","campaign_id":"23843461555820022","campaign":"(X\/\u5e38\u614b\u671f\u9593\/CPA\/\u896a\u5b50_\u820a\u5ba2_CPA)\u896a\u5b50_\u820a\u5ba2_CPA#[CUE00008182][Facebook][Footer][X][X][X][Banner][CPA]","adset_id":"23843461571790022","adset":"(RT\/\u52a0\u5165\u8cfc\u7269\u8eca\/X\/30days+member30days)cart30days+member30days_all#[X]","ad_id":"23843461710800022","ad":"(X\/\u6a5f\u80fd\/X\/Star\/\u6a5f\u80fd\u896a\/\u5f37\u8abf\u6a5f\u80fd\/HP\/\u55ae\u5716_\u9b54\u93e1_\u8f15\u9b06\u53bb\u5473)\u55ae\u5716_\u9664\u81ed\u7cfb\u5217_RT_190603#[Banner][1080x1080]","ad_type_id":1,"device_id":2,"day":"2020-01-01","action_values_offsite_conversion_fb_pixel_view_content_value":57800,"action_values_offsite_conversion_fb_pixel_purchase_value":12798,"action_values_offsite_conversion_fb_pixel_initiate_checkout_value":2460,"action_values_offsite_conversion_fb_pixel_add_to_cart_value":42180,"action_values_offsite_conversion_fb_pixel_add_payment_info_value":1910,"action_values_omni_view_content_value":57800,"action_values_omni_purchase_value":12798,"action_values_lead_value":0,"action_values_omni_initiated_checkout_value":2460,"action_values_omni_complete_registration_value":0,"action_values_omni_add_to_cart_value":42180,"action_values_add_payment_info_value":1910,"actions_post_reaction_value":4,"actions_offsite_conversion_fb_pixel_view_content_value":169,"actions_offsite_conversion_fb_pixel_purchase_value":8,"actions_offsite_conversion_fb_pixel_lead_value":8,"actions_offsite_conversion_fb_pixel_initiate_checkout_value":7,"actions_offsite_conversion_fb_pixel_complete_registration_value":6,"actions_offsite_conversion_fb_pixel_add_to_cart_value":144,"actions_offsite_conversion_fb_pixel_add_payment_info_value":6,"actions_link_click_value":87,"actions_landing_page_view_value":72,"actions_omni_view_content_value":169,"actions_omni_purchase_value":8,"actions_lead_value":8,"actions_omni_initiated_checkout_value":7,"actions_omni_complete_registration_value":6,"actions_omni_add_to_cart_value":144,"actions_add_payment_info_value":6,"actions_page_engagement_value":91,"actions_post_engagement_value":91,"clicks":163,"frequency":2.7854020000000004,"impressions":4956,"inline_link_clicks":87,"inline_post_engagement":91,"purchase_roas_omni_purchase_value":6.959217,"reach":3043,"spend":1852,"website_purchase_roas_offsite_conversion_fb_pixel_purchase_value":6.959217,"cost":1852,"conversions":8,"results":8}
{"media_account_uid":"1529888140439425","media_account":"Footer","campaign_id":"23843461555820022","campaign":"(X\/\u5e38\u614b\u671f\u9593\/CPA\/\u896a\u5b50_\u820a\u5ba2_CPA)\u896a\u5b50_\u820a\u5ba2_CPA#[CUE00008182][Facebook][Footer][X][X][X][Banner][CPA]","adset_id":"23843467721580022","adset":"(RT\/\u700f\u89bd\u7522\u54c1\u9801\/X\/10days)viewcontent10days_all#[X]","ad_id":"23843467721590022","ad":"(X\/\u6a5f\u80fd\/X\/Star\/\u6a5f\u80fd\u896a\/\u5f37\u8abf\u6a5f\u80fd\/HP\/\u55ae\u5716_\u9b54\u93e1_\u8f15\u9b06\u53bb\u5473)\u55ae\u5716_\u9664\u81ed\u7cfb\u5217_RT_190603#[Banner][1080x1080]","ad_type_id":1,"device_id":1,"day":"2020-01-01","clicks":0,"frequency":1.071429,"impressions":15,"inline_link_clicks":0,"inline_post_engagement":0,"reach":14,"spend":5,"cost":5}
```

### Executing program

golang
```
cd golang
go run main.go
```

nodejs
```
cd nodejs
node main.js
```

## Benchmark

### Golang

| File | Size | Output |
| ---- | ---- | ------ |
| small.log | 9K | 4K |
| middle.log | 80MB | 900K |
| large.log | 500MB | 5.4MB |

### Nodejs

| File | Size | Output |
| ---- | ---- | ------ |
| small.log | 9K | 1.3K |
| middle.log | 80MB | 135MB |
| large.log | 500MB | 886MB |

## Authors

* [CHING WEN, WANG](http://linkedin.com/in/cingwun0817)

## License

This project is licensed under the [MIT] License

## Acknowledgments

* [golang-parquet](https://github.com/xitongsys/parquet-go)
* [nodejs-parquet](https://github.com/ironSource/parquetjs)
