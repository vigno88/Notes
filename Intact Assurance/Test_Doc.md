# Test Learning

List tasks :
- modify spring boot code
- modify web code
- using jenkins
- using maven
- using github
- using scrum
- 
## General Info

- [Jenkins Url](https://stha38e56:444/view/Billing/)

  - Once in the flow, we can see a list of build
    - BUILD_ENDTOEND_GWBC-BDI = gobrio (old system that we want to replace by Guidewire)
    - BUILD_ENDTOEND_GWBC = Guidewire
    - D1 <- D2 <- D3
- To see the result of the test int the ITA gui, click on **Results URL**
- Do not note in the suite follow up anything related to **Nouvelle Affaire**
- Doc [ITA](https://apps.iad.ca.inet/sites/adsth/ITA/default.aspx )
  
## Tasks to do in the morning

[3.Scheduled Tests [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/)

Always verify if the fail already happened in the past
BDI in priority since other test depends of them later in the day

> **BDI** username : look at the execution list 
>         password :dadev999

- [x] [BUILD_ENDTOEND_GWBC-GW_D1_FLOW](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/BUILD_ENDTOEND_GWBC-GW_D1_FLOW/)
  > passer
- [x] [BUILD_ENDTOEND_GWBC-GW_D2_FLOW](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/BUILD_ENDTOEND_GWBC-GW_D2_FLOW/)
- [x] [BUILD_ENDTOEND_GWBC-GW_D3_FLOW](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/BUILD_ENDTOEND_GWBC-GW_D3_FLOW/)
- [x] [BUILD_ENDTOEND_GWBC-BDI_D1_FLOW [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/BUILD_ENDTOEND_GWBC-BDI_D1_FLOW/)
- [x] [BUILD_ENDTOEND_GWBC-BDI_D2_FLOW [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/BUILD_ENDTOEND_GWBC-BDI_D2_FLOW/)
- [x] [BUILD_ENDTOEND_GWBC-BDI_D3_FLOW [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/BUILD_ENDTOEND_GWBC-BDI_D3_FLOW/)
- [x] [BUILD_ENDTOEND_GWBC-BDI_FLOW [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/BUILD_ENDTOEND_GWBC-BDI_FLOW/)
- [ ] Look at the ping in the most recent build in [TEST_Integration [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/TEST_Integration/)
  > encore beaucoup de BIA et  BAA 
- [x] [TEST_SERVICES_JMeter [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/TEST_SERVICES_JMeter/)
  > BPTP_90098_TestProcessPolicy - Success:0, Fail:1 (Nicolas ma dit normal)
- [x] [TESTS_ALL_GWBC_FLOW [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/TESTS_ALL_GWBC_FLOW/)
  > 674-ProcessIncomingBankPaymentsFile
  > 107-CreateHistoryEventManualPaymentRequestandCancellation_Flow_D2

  > testCreateMultiplePaymentEntryBEL_BNA_ON_EN

  *withdrawal day*
  testPostdatedPaymentRemoveBEL_BEL_AB_EN
  testPostdatedPaymentRemoveBEL_BNA_ON_EN
  testApplySuspensePaymentBEL_BEL_AB_EN
  testApplySuspensePaymentBEL_BNA_ON_EN 
  testMoveSuspensePaymentBEL_BEL_AB_EN 
  testMoveSuspensePaymentBEL_BNA_ON_EN 
  testDistributeUnappliedFundsBEL_BEL_AB_EN 
   testDistributeUnappliedFundsBEL_BNA_ON_EN

    *Defect*
   testChangePaymentInfoIQC_NIL_QC_EN 


- [x] [TESTS_billing-BP-Clock [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/TESTS_billing-BP-Clock/)
  >  fail BP assert field withdrawalDay 
- [x] [TESTS_billing-delinquency_part1 [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/TESTS_billing-delinquency_part1/)
- [x] [TESTS_billing-delinquency_part2 [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/TESTS_billing-delinquency_part2/)
  >  1621-NonResponsiveDelinquency_TestCase7_CL (meme qu'hier)
- [x] [TESTS_Billing-statement [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/TESTS_Billing-statement/)
  > Daily batch on fail donc "normal" que D4 fail (pas de rapport pour les batch)
- [x] [TESTS_WORKFLOW [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/TESTS_WORKFLOW/)
    > magnifique
---

## How to fill Testsuit follow-up

- Test Suite
  - Name of the suite in Jenkins
- Test failing
  - Can be found in the HTML report
- Branch
  - Can be found in the HTML report
- Date
  - suite dashboard
- Environment
  - **AppConfigEnv** in the dashboard of the test report
- Report
  - The url of the test report
- Execution host
  - **Hostname** in the test report
- Screen def
  - In the test report, second menu from the top, in the test scrollHonolulu folder as been deletedmenu, find the test that failed. In the description of the test, the screendef can be found in the step before the one that failed.
- Test fails since:
  - Date of the failure
- Detail:
  - Conclusion of the investigation

## Test avec proxy

603-DesktopActivities Didn't fail
1008-SearchActivities Didn't fail

To check next time :

 356-ActivityEnhAcc 

## Task


- [x] sans proxy chrome
- [X] sans proxy ie
- [x] 1 start-end chrome
- [x] 1 start-end ie
- [x] 2 start-end chrome
- [x] 2 start-end ie
- [x] 1 per-keyword chrome
- [x] 1 per-keyword ie
- [x] 2 per-keyword chrome
- [x] 2 per-keyword ie
- [x] on error ie
- [x] on error chrome


https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/TESTS_ALL_GWBC_FLOW/373/TESTS_5fALL_5fGWBC_5fFLOW-373_5fHTML-Report/

https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/TESTS_ALL_GWBC_FLOW/353/TESTS_5fALL_5fGWBC_5fFLOW-353_5fHTML-Report/ItaFlowReport.html