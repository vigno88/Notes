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


- [x] [BUILD_ENDTOEND_GWBC-BDI_D1_FLOW [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/BUILD_ENDTOEND_GWBC-BDI_D1_FLOW/)
  > <span style="color:green">Didn't run today</span>
- [x] [BUILD_ENDTOEND_GWBC-BDI_D2_FLOW [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/BUILD_ENDTOEND_GWBC-BDI_D2_FLOW/)
  > <span style="color:green">Didn't run today</span>
- [ ] [BUILD_ENDTOEND_GWBC-BDI_D3_FLOW [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/BUILD_ENDTOEND_GWBC-BDI_D3_FLOW/)
  > <span style="color:green">Will run at 9AM</span>
- [x] [BUILD_ENDTOEND_GWBC-BDI_FLOW [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/BUILD_ENDTOEND_GWBC-BDI_FLOW/)
- [x] Look at the ping in the most recent build in [TEST_Integration [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/TEST_Integration/)
  > <span style="color:red">25 fails</span>
  > - 25 failures in TestNonPROD
- [x] [TEST_SERVICES_JMeter [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/TEST_SERVICES_JMeter/)
  > <span style="color:green">BAS_102_TestPolicyTransaction - Success:5, Fail:63</span> - we are a 31
  > <span style="color:green">BIA_86705_TestBillingInvoice - Success:55, Fail:6</span> - Fred
  > <span style="color:green">BPA_86537_RetrievePostdatedPayments - Success:22, Fail:5</span> - Antho  
- [x] [TESTS_ALL_GWBC_FLOW [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/TESTS_ALL_GWBC_FLOW/)
  > <span style="color:red">7 fails</span>
  >  674-ProcessIncomingBankPaymentsFile  <span style="color:green">To be launched</span>
  > 355-TransactionHistoryPolicyChangeCreditDistributionRefundSent_CL <span style="color:green">To be launched</span>
  > 	testMovePaymentBEL_BEL_AB_EN, phil on it
- [x] [TESTS_billing-BP-Clock [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/TESTS_billing-BP-Clock/) 
- [x] [TESTS_billing-delinquency_part1 [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/TESTS_billing-delinquency_part1/)
  > - 1123-DelinquencyCancelled_CL <span style="color:green">Adler on it</span>
- [x] [TESTS_billing-delinquency_part2 [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/TESTS_billing-delinquency_part2/)
  > -  1621-DelinquencyCancellationRenewalNonResponsive_CL <span style="color:green">To be launched</span> 

- [x] [TESTS_Billing-statement [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/TESTS_Billing-statement/) getDate script - 1133 et daily batch
  >  8_ReinstatementPaymentReversal 
- [x] [TESTS_WORKFLOW [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/TESTS_WORKFLOW/)
  > - WorkflowEFTRefund_Master_CL  <span style="color:red">Where is the test in ITL?</span>

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