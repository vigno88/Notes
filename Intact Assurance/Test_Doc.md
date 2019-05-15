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

- [ ] [BUILD_ENDTOEND_GWBC-BDI_D1_FLOW [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/BUILD_ENDTOEND_GWBC-BDI_D1_FLOW/)
  > <span style="color:red">NO Success, everything failed<span>
- [X] [BUILD_ENDTOEND_GWBC-BDI_D2_FLOW [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/BUILD_ENDTOEND_GWBC-BDI_D2_FLOW/)
- [ ] [BUILD_ENDTOEND_GWBC-BDI_D3_FLOW [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/BUILD_ENDTOEND_GWBC-BDI_D3_FLOW/)
  > <span style="color:red">NO Success, everything failed<span>
- [X] [BUILD_ENDTOEND_GWBC-BDI_FLOW [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/BUILD_ENDTOEND_GWBC-BDI_FLOW/)
- [x] Look at the ping in the most recent build in [TEST_Integration [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/TEST_Integration/)
  > 16 fail in TestNonPROD
  > 1 Maven Error
- [X] [TEST_SERVICES_JMeter [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/TEST_SERVICES_JMeter/)
- [ ] [TESTS_ALL_GWBC_FLOW [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/TESTS_ALL_GWBC_FLOW/)
  > <span style="color:red">36 fails<span>
  > - 1008-SearchActivities - <span style="color:green"> Launched & SUCCESS<span>
  > - 603-DesktopActivities - <span style="color:green"> Launched & SUCCESSS<span>
  > - 83742-CommissionPayment_CL  **as usual**
  > - 22-ManageAccountRefundsAuto_CL - <span style="color:green"> Launched & SUCCESS<span>
  > - Billing_BP_P1 **31 fails** - Normal (demander antoine pour resoudre le probleme)
  > - testVoidRefundIQC_NIL_QC_FR_SEQ2 <span style="color:red"> (IE) <span>  
- [ ] [TESTS_billing-BP-Clock [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/TESTS_billing-BP-Clock/)
  > <span style="color:red">1 missing<span> ->  testExternalCollectionsIQC_NIL_QC_FR_CLOCK 
- [X] [TESTS_billing-delinquency_part1 [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/TESTS_billing-delinquency_part1/)
- [X] [TESTS_billing-delinquency_part2 [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/TESTS_billing-delinquency_part2/)
  > 
- [X] [TESTS_Billing-statement [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/TESTS_Billing-statement/)
- [ ] [TESTS_WORKFLOW [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/TESTS_WORKFLOW/)
  > <span style="color:red">1 fail<span>
  > - WorkflowEFTRefund_Master_CL <span style="color:green">Launched & Success<span>

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