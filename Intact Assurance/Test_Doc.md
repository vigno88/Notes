# Test Learning

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


- [ ] [BUILD_ENDTOEND_GWBC-BDI_D1_FLOW [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/BUILD_ENDTOEND_GWBC-BDI_D1_FLOW/)
  > <span style="color:red">3 fails</span>
  > -  355-TransactionHistoryQcReplacingPolicy_BDI_D1 
  > -  178-EndToEndBelairQCAuto25N1NewBusiness_BDI_D1
  > - 178-EndToEndBelairBCAutoA1NewBusiness_BDI_D1
- [x] [BUILD_ENDTOEND_GWBC-BDI_D2_FLOW [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/BUILD_ENDTOEND_GWBC-BDI_D2_FLOW/)
  > <span style="color:red">Didn't run today<span>
- [x] [BUILD_ENDTOEND_GWBC-BDI_D3_FLOW [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/BUILD_ENDTOEND_GWBC-BDI_D3_FLOW/)
  > <span style="color:red">2 fails</span>
  >  178-EndToEndBelairABAutoB5Reinstatement_BDI_D3 
  > 178-EndToEndBelairBCAutoB5Reinstatement_BDI_D3
- [x] [BUILD_ENDTOEND_GWBC-BDI_FLOW [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/BUILD_ENDTOEND_GWBC-BDI_FLOW/)
- [ ] Look at the ping in the most recent build in [TEST_Integration [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/TEST_Integration/)
  > <span style="color:red">23 fails</span>
  > - 22 failures in TestNonPROD
  > - 1 fail in TestPingAllGWBC
- [x] [TEST_SERVICES_JMeter [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/TEST_SERVICES_JMeter/)
  > <span style="color:red">1 fail BSA<span>
- [x] [TESTS_ALL_GWBC_FLOW [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/TESTS_ALL_GWBC_FLOW/)
  > <span style="color:red">34 fails - 7 missing<span>
  > -  356-ActivityEnhAcc <span style="color:green">Pass<span>
  > -  603-DesktopActivities <span style="color:green">Pass<span>
  > -  83742-CommissionPayment_CL **Gab** 
  > -  SaveGwbcLogs <span style="color:green">To be launched<span> ??
  > -  testNegativeWriteOffReversalBEL_BEL_AB_EN<span style="color:red"> All languages<span> <span style="color:green">To be launched<span>
  > - testCompleteNegativeWriteOffRejectionBEL_BEL_AB_EN<span style="color:red"> All languages<span> <span style="color:green">To be launched<span>
  > -  testReverseWriteOffBEL_BEL_AB_EN <span style="color:red"> All languages<span> <span style="color:green">To be launched<span>
  > - testCompleteWriteOffRejectionBEL_BEL_AB_EN <span style="color:red"> All languages<span> <span style="color:green">To be launched<span>
  > - testVoidRefundBEL_BEL_AB_EN_SEQ1 <span style="color:red"> All languages<span> <span style="color:green">To be launched<span>
  > -  testEditRefundBEL_BEL_AB_EN <span style="color:red"> All languages<span> <span style="color:green">To be launched<span>
  > -  testRescindRefundBEL_BEL_AB_EN <span style="color:red"> All languages<span> <span style="color:green">To be launched<span>
  > - All of Billing_BP_P2 <span style="color:green">To be launched<span>
  > - <span style="color:blue">testCompleteNegativeWriteOffRejectionIQC_NIL_QC_FR <span>
  > - <span style="color:blue"> testCompleteTroubleTicketBEL_BNA_ON_EN <span>
  > - <span style="color:blue"> testEditRefundIQC_NIL_QC_FR <span>
  > - <span style="color:blue">  testRescindRefundIQC_NIL_QC_EN  <span>
  > - <span style="color:blue">  testRescindRefundIQC_NIL_QC_FR  <span>
  > - <span style="color:blue"> testRejectRefundBEL_BEL_AB_EN<span>
  > - <span style="color:blue">  testRejectRefundIQC_NIL_QC_EN <span>
  > - <span style="color:blue"> testRejectRefundIQC_NIL_QC_FR  <span>
- [ ] [TESTS_billing-BP-Clock [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/TESTS_billing-BP-Clock/)
    > <span style="color:red"> 1 missing<span>
    > - <span style="color:blue">testExternalCollectionsIQC_NIL_QC_FR_CLOCK- To be launched</span>   
- [x] [TESTS_billing-delinquency_part1 [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/TESTS_billing-delinquency_part1/)
- [x] [TESTS_billing-delinquency_part2 [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/TESTS_billing-delinquency_part2/)

- [x] [TESTS_Billing-statement [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/TESTS_Billing-statement/) getDate script - 1133 et daily batch
  > <span style="color:red"> 5 fails<span>
  > -  3_NewBusinessRewriteCancellation_D2 <span style="color:green">To be launched<span>
  > -  3_NewBusinessRewriteCancellation_D3 <span style="color:green">To be launched<span>
  > -   8_ReinstatementPaymentReversal_A_D2 <span style="color:green">To be launched<span>
  > -  8_ReinstatementPaymentReversal_A_D3 <span style="color:green">To be launched<span>
  > -  8_ReinstatementPaymentReversal_A_D4 <span style="color:green">To be launched<span>
- [x] [TESTS_WORKFLOW [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/TESTS_WORKFLOW/)

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
  - In the test report, second menu from the top, in the test scrolldown menu, find the test that failed. In the description of the test, the screendef can be found in the step before the one that failed.
- Test fails since:
  - Date of the failure
- Detail:
  - Conclusion of the investigation