# poc-go-file-splitter

```
go test -bench=. -benchmem -memprofile memprofile.out -cpuprofile profile.out                                                                                     
```

```
go tool pprof profile.out
```

doc
```
Saved: 5f5bcf47-9fb9-4186-88ad-7c7518a9dd49.csv after 25.032970335s
Saved: 78f4d5f7-a98b-466a-8541-435be4710de0.csv after 3m39.174302229s
Saved: 8d69258c-a558-4a35-aaa9-4e6c5809999f.csv after 3m45.589299391s
Saved: dc1ec63d-5611-4e29-b501-1e89dc45cc7a.csv after 3m47.243862635s
Saved: 916c33b8-4e0a-47c3-9b0f-447c3eb02d67.csv after 3m52.322252401s
Saved: 779fed10-bcf4-43d4-a0dc-a4300d4618d1.csv after 4m6.236435071s
Saved: 4a047d04-55f9-4609-b42b-5fb987e81cb0.csv after 4m11.431136889s
Saved: 57174303-b40e-4b17-83d6-35eb9b52d4f5.csv after 4m11.859277495s
Saved: fafd740b-5fdb-459e-9d89-fa294540529a.csv after 4m14.829979906s
Saved: b8fcfb39-55c5-4bff-8828-672585f9c3cd.csv after 4m21.47974764s
Saved: 40ded071-b673-4844-bb44-0699cc9da62b.csv after 4m26.536646288s
Saved: d72f796f-0b37-46ee-9b3b-71f8366412c0.csv after 4m27.885321068s
Saved: 0451f9be-3603-425a-b9ef-c2775f53c3be.csv after 4m29.210837177s
Saved: 0a1d326c-c17b-4a07-9bd7-5c1806278a5d.csv after 4m26.993844324s
Saved: 65771c6e-4948-42a5-94ae-c72bc882fa88.csv after 4m30.761761067s
Saved: 0ce2a691-2dc9-41d8-9916-ef86ab902905.csv after 4m30.416348109s
Saved: eda08cd8-b319-4574-81e6-1e261a58c0e8.csv after 4m30.410294976s
Saved: b69b6c43-7777-45d5-9435-14ce1e836b8c.csv after 4m31.781667213s
Saved: 9fcd0610-c1a1-4ce7-bb3d-f81cd00362b3.csv after 4m35.105286519s
Saved: 53b92b17-4d23-4e1a-9149-ee748606a6db.csv after 4m35.703986621s
Saved: c2eeab56-ad90-4f1f-b789-cdfb5ec943ab.csv after 4m35.416248687s
Saved: 586a4d12-e062-4173-9a01-6c2eb69aefc9.csv after 4m41.250205346s
Saved: db02bdda-2669-47d3-a534-623a51bd8739.csv after 4m41.091748655s
Saved: f7c4b207-1c14-4ec4-8c5b-e1d0aa744844.csv after 4m41.098511509s
Saved: ee78fbf2-1c63-46d2-bcef-f9eb88f041c0.csv after 4m41.530377759s
Saved: 39ac576e-178f-4210-89a7-0a208752f72e.csv after 4m43.776751341s
Saved: fd6ddf9c-b9f5-4e47-b4dd-93d814bc6970.csv after 4m44.451379559s
Saved: a05757d0-5f36-4991-9c95-57708f2b0d0f.csv after 4m47.024693592s
Saved: 162ccba2-7bc7-4a05-a858-b78d51fd6304.csv after 4m51.58803875s
Saved: 3e315092-d6cc-4603-8c71-070228402588.csv after 4m54.687765588s
Saved: c2110d7e-e0e2-407e-a55f-1d46ee9b7119.csv after 4m53.906724445s
Saved: 1f51a726-2101-41d0-8fed-e90a519e1f46.csv after 4m55.437664967s
Saved: cf45e4d4-5352-4415-8065-dd30ca315a46.csv after 4m59.704051484s
Saved: 4c09ce98-c0ed-4da9-88d6-e230894dbbad.csv after 5m1.110881498s
Saved: d730deaa-5dea-49f0-8314-8f8862c497d6.csv after 5m0.49411317s
Saved: 3f6e4484-6cbb-435a-ad13-b735c5835047.csv after 5m0.936835983s
Saved: 6fad8933-b7f5-4c2d-bb63-a48b5aa8cda6.csv after 5m0.772922515s
Saved: 2b381ad1-8a90-4e80-bd91-37f3d6ae0f35.csv after 5m2.807057266s
Saved: 2f8b6246-dd35-4d21-9563-ea355e56bdb3.csv after 5m4.3932713s
Saved: 3f1df5e3-e115-4e33-9d01-bec99a091fcc.csv after 5m4.733483155s
Saved: 110cb94b-2981-4da3-a66e-cfcd160161c2.csv after 5m6.93134391s
Saved: f892fc8e-15d7-4a29-bc29-9c38f5534a21.csv after 5m6.095814018s
Saved: 68e9f71b-657f-455f-a2b6-d7d0966f4d18.csv after 5m8.992365289s
Saved: 037ea542-2665-4748-b475-4ce91bfd70c9.csv after 5m9.206546213s
Saved: db041247-a806-47e4-8a61-5971eeed0c3c.csv after 5m10.572171663s
Saved: a24db24f-738c-4fbe-ac57-c35ef80f815d.csv after 5m14.256691869s
Saved: 65c12435-1f06-4c16-ba49-aea7fe1acd71.csv after 5m16.181240856s
Saved: 5cbf91cb-2a2a-4214-aeaa-977fce1df7b4.csv after 5m24.015613662s
Saved: 19b829d3-ef0a-4e47-802a-6cc0978b8eee.csv after 5m23.556919764s
Saved: 5de6872b-e4c1-4308-8fa8-2c502fb0053a.csv after 5m27.022191503s
Saved: bbd71f7c-b726-48b7-9dad-b997df2b220a.csv after 5m29.246793387s
Saved: c43e95d7-7c94-4977-9cba-4786b6bcb235.csv after 5m36.299222482s
Saved: a043f7cc-9f81-426a-aed0-c1afefa76969.csv after 5m36.69009383s
Saved: e86d6814-5b96-4c3a-8074-8833c6d1af17.csv after 5m37.097111511s
Saved: c18d8b38-daeb-4674-a406-36ac4e973570.csv after 5m36.209643516s
Saved: 530a1465-f4a2-48aa-ab5b-4be54634da11.csv after 5m38.248598094s
Saved: 3c351c6d-2481-485e-b5b3-416ddf559898.csv after 5m40.23872763s
Saved: 23c1e6a8-3f74-4d1f-bbec-ab32d444ef14.csv after 5m41.528086215s
Saved: 9ffd187a-c2b9-4013-b678-e904b8fccf2e.csv after 5m44.075575214s
Saved: 5416fb58-fba7-4277-a37b-151f623160e4.csv after 5m47.337371883s
Saved: 705613e4-96d8-4327-8733-3425ab3a249c.csv after 5m48.587435149s
Saved: e080b4ea-760c-4ff3-8956-89dd3aff10cf.csv after 5m48.522310125s
Saved: 5265086e-60b3-4a1b-9ec2-db7e5dd1ad40.csv after 5m47.632718242s
Saved: 5583917a-86e8-4893-a2a3-4d4b01526d94.csv after 5m48.970681914s
Saved: 8dbc5b9f-2298-452b-8fe3-31e96c6c3233.csv after 5m49.161921963s
Saved: ae20df0d-fe90-49c9-84c2-c62e01733253.csv after 5m50.465273937s
Saved: d4e3f662-b404-4c67-8ede-2d81ae86f13e.csv after 5m53.834525752s
Saved: 7fc9558e-4513-463c-8dd6-67557ee392d1.csv after 5m55.73579166s
Saved: 1382cb85-c9f7-453d-bd7b-54f134a860d2.csv after 5m56.539652753s
Saved: 93f2c37c-1fbf-447e-93b2-0cefe72f8af2.csv after 5m58.109273448s
Saved: 762ac849-8598-43df-8eb0-f832b4027a6c.csv after 5m58.363792128s
Saved: 3b6f3eaa-b5d0-4e53-b0fd-3f0ce2097975.csv after 6m1.117577948s
Saved: 463f2c92-c5b9-4e6d-8aba-cd5c6b97f766.csv after 6m1.104300359s
Saved: dfccb558-7ffb-46ca-81ce-ae84ee2ebd2b.csv after 6m1.725160728s
Saved: 78fd4a07-d47c-400d-b42f-bd4d3eb70dd2.csv after 6m1.436435567s
Saved: 0fd6d9be-3409-49aa-aa7c-3321543c6db0.csv after 6m3.822759327s
Saved: 8aae3ff1-fc02-4d1d-aa7f-6cd3938b3914.csv after 6m2.378637616s
Saved: 3feba037-2953-49d6-b257-4657dedcc87c.csv after 6m4.647498327s
Saved: be1e2b8a-ebf4-4e39-b9dc-31ba82972ebb.csv after 6m3.86254586s
Saved: 6013424c-d023-49f9-b3e3-0c1bdc97e0b8.csv after 6m5.457290319s
Saved: 43c5a8a0-8de0-4e96-9f26-4c0f60a61f0b.csv after 6m5.244196804s
Saved: 979773a1-0250-47dc-b353-2924d8624f16.csv after 6m5.951635091s
Saved: 559b0a99-1ae9-45a9-a16b-08c19523e175.csv after 6m5.827525071s
Saved: f69b6c8e-4051-4e40-996f-bac3c6e9db97.csv after 6m6.499223224s
Saved: f8dd0055-d3c1-4122-82ea-09205d2aa9b6.csv after 6m7.41329837s
Saved: baadedaa-47aa-4045-a7de-0dca191d05a2.csv after 6m5.283049004s
Saved: 8ce47b24-ecbc-414a-8cde-71290eb6fa47.csv after 6m8.112619428s
Saved: 3fd0b0c2-2c10-4509-b701-ba278f7de6f7.csv after 6m7.922593248s
```