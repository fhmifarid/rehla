# Rehla Platform

## Detailed Product & Engineering Specification — v2.0

---

# 1. الهدف من هذه الوثيقة

هذه الوثيقة هي المرجع الوظيفي والتقني الأساسي لبناء:

* تطبيق العميل Flutter.
* لوحة التحكم Next.js.
* Backend بلغة Go.
* PostgreSQL.
* نظام الطلبات.
* نظام الدفع بالتحويل البنكي.
* نظام المحفظة.
* إدارة الباقات والخدمات.
* إدارة العملاء والمسافرين.
* التشغيل والرحلات.
* الدعم.
* التسويق.
* الإشعارات.
* التقارير.
* الأتمتة.
* الأمان.
* الضبط الكامل للنظام.

الهدف أن يستطيع أي فريق Backend أو Frontend قراءة القسم الخاص به ومعرفة:

1. ما الصفحات المطلوبة.
2. ما الذي يظهر داخل كل صفحة.
3. ما العمليات المتاحة.
4. ما الحقول المطلوبة.
5. ما حالات البيانات.
6. ما القيود.
7. ما الصلاحيات.
8. ما السلوك المطلوب من Go Backend.
9. ما الذي يجب تسجيله في Audit Log.

---

# 2. قواعد عامة لجميع أقسام لوحة التحكم

كل صفحة List في النظام يجب أن تدعم - متى كان ذلك منطقيًا:

* Pagination.
* Search.
* Filters.
* Sorting.
* Saved Views.
* Column visibility.
* Date range.
* Export.
* Refresh.
* Bulk selection للعمليات الآمنة.
* Empty State.
* Loading State.
* Error State.
* Permission-aware actions.

كل سجل يجب أن يدعم:

* Created At.
* Updated At.
* Created By.
* Updated By.
* Status.
* Activity Timeline عند الحاجة.
* Audit Trail للعمليات الحساسة.

الحذف لا يستخدم دائمًا `DELETE` حقيقي.

يجب التمييز بين:

* Delete.
* Archive.
* Disable.
* Cancel.
* Reverse.

خصوصًا في البيانات المالية والتشغيلية.

---

# 3. هيكل Sidebar النهائي

```text
Home
├── Dashboard
├── Activity
└── My Tasks

Commerce
├── Orders
├── Customers
├── Packages
├── Products & Add-ons
├── Categories
├── Destinations
├── Package Types
├── Pricing
├── Promotions
└── Coupons

Travel Operations
├── Departures
├── Availability
├── Reservations
├── Travelers
├── Documents
├── Pickup Points
├── Waitlists
├── Manifests
└── Operations Tasks

Finance
├── Overview
├── Payments
├── Payment Requests
├── Payment Proofs
├── Bank Transactions
├── Reconciliation
├── Unmatched Transfers
├── Wallets
├── Wallet Top-ups
├── Wallet Transactions
├── Wallet Holds
├── Wallet Adjustments
├── Customer Credits
├── Refunds
├── Financial Ledger
└── Reports

Customer Experience
├── Support Inbox
├── Tickets
├── Chat
├── Reviews
└── Notifications

Growth
├── Customer Segments
├── Campaigns
├── Loyalty
└── Referrals

Partners
├── Vendors
├── Suppliers
├── Contracts
└── Supplier Services

Content
├── Home Builder
├── Pages
├── Banners
├── FAQs
├── Policies
└── Media Library

Analytics
├── Overview
├── Sales
├── Orders
├── Customers
├── Packages
├── Operations
├── Finance
├── Wallet
├── Marketing
└── Custom Reports

Automation
├── Workflows
├── Approvals
├── Scheduled Jobs
└── Job Monitor

Security
├── Staff
├── Teams
├── Roles
├── Permissions
├── Sessions
├── Audit Logs
└── Risk Flags

System
├── Integrations
├── Webhooks
├── API Keys
├── Imports
├── Exports
├── API Logs
└── Feature Flags

Settings
```

---

# 4. Dashboard

## 4.1 الصفحة الرئيسية

المسار:

```text
/dashboard
```

يجب أن تحتوي على Dashboard ديناميكية حسب Role.

---

## 4.2 KPIs التجارة

بطاقات:

* إجمالي الطلبات.
* الطلبات الجديدة.
* الطلبات المدفوعة.
* الطلبات غير المدفوعة.
* الطلبات المدفوعة جزئيًا.
* الطلبات الملغاة.
* الطلبات المكتملة.
* متوسط قيمة الطلب.
* إجمالي المبيعات.
* صافي المبيعات.
* المبلغ المستحق من العملاء.

---

## 4.3 KPIs المالية

* التحويلات قيد المراجعة.
* إيصالات تحتاج مراجعة.
* التحويلات غير المطابقة.
* Refunds معلقة.
* Refunds مكتملة.
* الأموال المستلمة اليوم.
* صافي التحصيل.
* Wallet Liability.
* الرصيد المعلق بالمحافظ.
* Wallet Top-ups قيد المراجعة.

---

## 4.4 KPIs التشغيل

* الرحلات القادمة اليوم.
* الرحلات خلال 7 أيام.
* الرحلات ممتلئة.
* الرحلات قاربت على الامتلاء.
* عدد المقاعد المتاحة.
* Waitlist.
* مسافرون لديهم مستندات ناقصة.
* مستندات منتهية أو قريبة الانتهاء.

---

## 4.5 Charts

* Revenue by Day.
* Orders by Day.
* Orders by Status.
* Payment Method Distribution.
* Wallet vs Bank Payments.
* Top Packages.
* Top Destinations.
* Cancellation Rate.
* Refund Rate.
* New vs Returning Customers.

---

## 4.6 Recent Activity

Timeline لحظية مثل:

```text
Order #12345 created
Payment proof submitted
Wallet top-up approved
Package updated
Refund completed
New customer registered
New support ticket
```

---

## 4.7 تخصيص Dashboard

الموظف يستطيع:

* إضافة Widget.
* إزالة Widget.
* تغيير الحجم.
* تغيير الموقع.
* تحديد الفترة.
* حفظ Layout.
* Reset Layout.

ويحفظ Layout لكل Staff User.

---

# 5. Orders

هذا القسم يجب أن يكون من أكبر أقسام النظام.

---

# 5.1 صفحة جميع الطلبات

```text
/orders
```

Columns:

* Order Number.
* Customer.
* Package.
* Departure.
* Travelers Count.
* Total.
* Paid.
* Balance Due.
* Payment Status.
* Order Status.
* Service Status.
* Refund Status.
* Risk.
* Priority.
* Assigned To.
* Created At.

---

## 5.2 Search

يبحث بواسطة:

* Order Number.
* Customer Name.
* Phone.
* Email.
* Payment Reference.
* Bank Reference.
* Traveler Name.
* Passport Number.

---

## 5.3 Filters

* Order Status.
* Payment Status.
* Service Status.
* Refund Status.
* Cancellation Status.
* Risk Status.
* Package.
* Destination.
* Departure.
* Assigned Staff.
* Tags.
* Date.
* Travel Date.
* Amount Range.
* Has Wallet Payment.
* Has Bank Payment.
* Has Pending Documents.

---

# 5.4 Saved Views

أمثلة:

* New Orders.
* My Orders.
* Awaiting Payment.
* Payment Review.
* Partially Paid.
* Paid But Unconfirmed.
* Traveling Tomorrow.
* Missing Documents.
* Cancellation Requests.
* Refund Pending.
* High Value.
* VIP.
* On Hold.

---

# 5.5 إنشاء طلب يدوي

```text
/orders/new
```

## القسم الأول — العميل

حقول:

* Customer Search.
* Create New Customer.
* Customer Name.
* Phone.
* Email.

---

## القسم الثاني — الباقة

* Package.
* Package Variant.
* Departure.
* Dates.
* Pickup Point.

---

## القسم الثالث — المسافرون

* اختيار مسافرين محفوظين.
* إضافة مسافر جديد.
* عدد المسافرين.

---

## القسم الرابع — Products & Add-ons

يمكن اختيار:

* منتجات.
* خدمات إضافية.
* Upgrade.
* Bundle.

مع Quantity.

---

## القسم الخامس — Pricing

يظهر Backend calculation:

```text
Package                8,000
Add-ons                 500
Products                300
Discount               -500
Manual Adjustment       200
----------------------------
Total                  8,500
```

---

## القسم السادس — Discount

يدعم:

* Coupon.
* Automatic Promotion.
* Manual Percentage Discount.
* Manual Fixed Discount.

Manual Discount يحتاج:

* القيمة.
* السبب.
* ملاحظة داخلية.
* Approved By إذا تجاوز Threshold.

---

## القسم السابع — الدفع

خيارات:

* Wallet.
* Bank Transfer.
* Wallet + Bank.

ويجب عرض:

* Wallet Available.
* Wallet Selected.
* Remaining Bank Amount.

---

## Actions

* Save Draft.
* Calculate.
* Create Order.
* Create & Request Payment.
* Cancel.

---

# 5.6 Order Detail Page

```text
/orders/[id]
```

Header:

* Order Number.
* Customer.
* Created Date.
* Travel Date.
* Total.
* Paid.
* Balance.
* Status badges.

Actions:

* Edit.
* Duplicate.
* Hold.
* Cancel.
* Refund.
* Request Payment.
* Add Note.
* Send Message.
* Assign.
* Archive.

---

## Order Sections

### Summary

* Package.
* Departure.
* Traveler count.
* Dates.
* Pickup.

### Customer

* Name.
* Phone.
* Email.
* Customer Status.
* Customer Tags.

### Travelers

لكل مسافر:

* Name.
* Passport.
* Document Status.
* Verification.

### Items

* Package.
* Products.
* Add-ons.
* Adjustments.
* Discounts.

### Financial Summary

```text
Subtotal
Discount
Adjustments
Total
Paid
Refunded
Net Paid
Balance Due
```

### Payments

كل Payment Source:

* Wallet.
* Bank.
* Manual Adjustment.

### Cancellation

* Request.
* Policy.
* Reason.
* Refund calculation.

### Documents

* Complete.
* Missing.
* Rejected.
* Expired.

### Internal Notes

لا تظهر للعميل.

### Customer Messages

تظهر للعميل.

### Tags

Custom.

### Timeline

جميع الأحداث.

### Audit

التغييرات الحساسة.

---

# 5.7 تعديل الطلب

المسموح:

* تغيير Departure.
* تغيير Dates.
* إضافة/حذف Traveler.
* إضافة Product.
* حذف Product.
* إضافة Add-on.
* تغيير Quantity.
* تغيير Pickup.
* إضافة Discount.
* إضافة Adjustment.

بعد التعديل Backend يعيد حساب:

```text
Old Total
New Total
Already Paid
Balance Due
Refund Due
```

ويتم إنشاء Revision.

---

# 5.8 Order Revisions

كل تعديل يحتوي:

* Revision Number.
* Changed By.
* Reason.
* Before Snapshot.
* After Snapshot.
* Financial Impact.
* Timestamp.

---

# 6. Customers

---

# 6.1 قائمة العملاء

```text
/customers
```

Columns:

* Name.
* Phone.
* Email.
* Status.
* Orders Count.
* Total Spend.
* Wallet Balance.
* Travelers Count.
* Customer Segment.
* Last Order.
* Registered At.

---

## Actions

* View.
* Edit.
* Suspend.
* Activate.
* Add Tag.
* Add Note.
* Create Order.
* Open Wallet.
* Message.
* Merge.

---

# 6.2 إضافة عميل

Fields:

* Full Name.
* Email.
* Phone.
* Password أو Send Setup Link.
* Status.
* Language.
* Tags.
* Internal Notes.

---

# 6.3 Customer Profile

Tabs:

```text
Overview
Orders
Travelers
Wallet
Payments
Refunds
Top-ups
Reviews
Support
Notes
Timeline
Security
```

---

# 6.4 Customer Overview

* Lifetime Spend.
* Orders.
* Completed Orders.
* Cancelled Orders.
* Refund Amount.
* Average Order Value.
* Wallet.
* Last Activity.
* Segment.
* Risk.

---

# 6.5 Customer Status

```text
ACTIVE
RESTRICTED
SUSPENDED
CLOSED
```

Suspend requires:

* Reason.
* Staff Note.

---

# 7. Travelers

---

# 7.1 Travelers List

Columns:

* Full Name.
* Customer.
* Passport Number.
* Nationality.
* Passport Expiry.
* Verification Status.
* Last Travel.
* Documents Status.

---

# 7.2 إضافة مسافر

Fields:

### Identity

* Full Name.
* First Name.
* Middle Name.
* Last Name.
* Date of Birth.
* Gender.
* Nationality.
* Relationship to Account Owner.

### Passport

* Passport Number.
* Issue Date.
* Expiration Date.
* Issuing Country.

### Files

* Passport Image.
* Personal Photo.

### Other

* Notes.
* Special Requirements.

---

# 7.3 Traveler Status

```text
ACTIVE
NEEDS_INFORMATION
DOCUMENTS_PENDING
VERIFIED
RESTRICTED
```

---

# 7.4 Document Verification

Actions:

* Approve.
* Reject.
* Request New Copy.
* Mark Expired.
* Add Note.

رفض المستند يحتاج سببًا.

---

# 8. Packages

---

# 8.1 قائمة الباقات

```text
/packages
```

Columns:

* Image.
* Package Name.
* Category.
* Destination.
* Type.
* Price.
* Discount.
* Departures.
* Status.
* Featured.
* Updated At.

Actions:

* View.
* Edit.
* Duplicate.
* Publish.
* Unpublish.
* Pause.
* Archive.
* Delete Draft.

---

# 8.2 Add Package

```text
/packages/new
```

يجب تقسيم النموذج إلى Tabs/Steps.

---

## Step 1 — Basic Information

Fields:

* Package Name *
* Slug
* Short Description *
* Full Description *
* Package Type *
* Category *
* Destination *
* Tags
* Internal Code / SKU
* Status

---

## Step 2 — Pricing

* Base Price *
* Currency
* Discount Type:

  * None
  * Percentage
  * Fixed
* Discount Value
* Promotional Price preview
* Minimum Travelers
* Maximum Travelers إذا كان عامًا

---

## Step 3 — Media

* Featured Image *
* Image Gallery
* 360° Image
* Video URL اختياري
* Alt Text لكل صورة
* Image Order

Actions:

* Upload.
* Reorder.
* Set Featured.
* Remove.

---

## Step 4 — Location

* Location Name.
* Latitude.
* Longitude.
* Map Picker.

---

## Step 5 — Highlights

Dynamic list:

```text
+ Add Highlight
```

كل عنصر:

* Title.
* Description اختياري.
* Order.

---

## Step 6 — Inclusions

Dynamic list مثل:

* Hotel.
* Transportation.
* Meals.
* Guide.

---

## Step 7 — Exclusions

Dynamic list.

---

## Step 8 — Pickup Points

اختيار من Pickup Points الحالية.

ويمكن إضافة:

* Additional Instructions.

---

## Step 9 — Policies

* Cancellation Policy.
* Refund Policy.
* Terms.
* Minimum Notice.
* Booking Deadline.

---

## Step 10 — Documents

تحديد المستندات المطلوبة:

* Passport.
* Photo.
* Visa.
* Custom Document.

لكل Requirement:

* Required / Optional.
* Expiry Validation.
* Minimum Validity Days.

---

## Step 11 — Add-ons

اختيار Add-ons المسموحة للباقة.

---

## Step 12 — SEO

* SEO Title.
* Meta Description.
* Keywords.
* Social Image.

---

## Step 13 — Publishing

* Draft.
* Publish Now.
* Schedule Publish.
* Scheduled Unpublish.
* Featured.

---

# 8.3 Package Status

```text
DRAFT
SCHEDULED
ACTIVE
PAUSED
SOLD_OUT
ARCHIVED
```

---

# 9. Products & Add-ons

هذا مثال على مستوى التفصيل المطلوب في جميع الأقسام.

---

# 9.1 Products List

```text
/products
```

Columns:

* Image.
* Name.
* Internal Code.
* Type.
* Category.
* Price.
* Stock/Availability.
* Eligible Packages.
* Status.
* Published.
* Updated At.

Search:

* Name.
* SKU.
* Category.

Filters:

* Type.
* Status.
* Published.
* Category.
* Package.

Actions:

* View.
* Edit.
* Duplicate.
* Publish.
* Unpublish.
* Archive.
* Delete Draft.

---

# 9.2 إضافة Product

```text
/products/new
```

## Basic Information

* Product Name *
* Internal Code / SKU
* Product Type *

Types:

```text
PHYSICAL_PRODUCT
SERVICE
ADD_ON
UPGRADE
BUNDLE
```

* Category *
* Short Description
* Full Description
* Tags

---

## Pricing

* Base Price *
* Currency
* Cost Price داخلي
* Compare At Price اختياري
* Discount Type
* Discount Value

يظهر:

* Selling Price.
* Cost.
* Estimated Margin.

---

## Media

* Featured Image.
* Gallery.
* Alt Text.
* Reorder.

---

## Availability

اختيار:

```text
Always Available
Limited Quantity
Departure-Based
Package-Based
```

إذا Limited:

* Available Quantity.
* Low Stock Threshold.

---

## Eligibility

يمكن تحديد:

* All Packages.
* Selected Packages.
* Selected Categories.
* Selected Destinations.
* Selected Departures.

---

## Purchase Rules

* Minimum Quantity.
* Maximum Quantity.
* Max Per Traveler.
* Max Per Order.
* Available Before Travel بـX ساعات/أيام.
* Available After Booking.
* Refundable.
* Non-refundable.
* Cancellation Deadline.

---

## Bundle Configuration

إذا Type = Bundle:

* Included Products.
* Included Quantities.
* Bundle Price.

---

## Tax / Accounting Metadata

حتى لو لم يستخدم النظام Tax حاليًا:

* Accounting Category.
* Revenue Category.
* Supplier اختياري.

---

## Publishing

* Draft.
* Active.
* Scheduled.
* Paused.
* Archived.

Fields:

* Publish At.
* Unpublish At.

---

# 9.3 Edit Product

يستخدم نفس النموذج.

يجب تسجيل:

* Before.
* After.
* Staff.
* Time.

تغيير السعر لا يغير الطلبات القديمة.

---

# 9.4 Delete Product

الحذف الحقيقي مسموح فقط إن كان:

* Draft.
* لم يستخدم في أي Order.

وإلا:

```text
Archive
```

---

# 10. Categories

---

## List

* Name.
* Slug.
* Image.
* Packages Count.
* Status.
* Order.

Actions:

* Add.
* Edit.
* Reorder.
* Disable.
* Archive.

---

## Add Category Fields

* Name.
* Slug.
* Description.
* Image.
* Icon.
* Parent Category اختياري.
* Display Order.
* Active.
* SEO Title.
* SEO Description.

---

# 11. Destinations

Fields:

* Destination Name.
* Country.
* City / Region.
* Description.
* Image.
* Gallery.
* Coordinates.
* Featured.
* Active.
* SEO.

ويعرض:

* Packages Count.
* Upcoming Departures.
* Orders.
* Revenue.

---

# 12. Package Types

مثلاً:

```text
Umrah
Tour
Hotel
Adventure
Religious
Family
```

Fields:

* Name.
* Description.
* Icon.
* Active.
* Order.

CRUD كامل.

---

# 13. Departures

---

# 13.1 Departures List

Columns:

* Departure Number.
* Package.
* Start Date.
* End Date.
* Capacity.
* Reserved.
* Confirmed.
* Available.
* Waitlist.
* Status.

---

# 13.2 Add Departure

Fields:

### Basic

* Package *
* Variant.
* Start Date & Time *
* End Date & Time *
* Booking Opens At.
* Booking Closes At.

### Capacity

* Total Capacity *
* Waitlist Enabled.
* Waitlist Capacity.

### Pricing Override

* Use Package Price.
* Custom Departure Price.
* Custom Discount.

### Operations

* Assigned Team.
* Internal Notes.
* Supplier.
* Pickup Points.

### Status

* Draft.
* Open.
* Closed.

---

# 13.3 Departure Detail

Tabs:

```text
Overview
Orders
Travelers
Availability
Waitlist
Documents
Payments
Pickup
Tasks
Suppliers
Timeline
```

---

# 14. Reservations

Reservation مؤقتة أثناء Checkout.

List يعرض:

* Reservation ID.
* Customer.
* Departure.
* Seats.
* Created.
* Expires.
* Status.

Statuses:

```text
ACTIVE
CONFIRMED
EXPIRED
RELEASED
CANCELLED
```

Admin يمكنه:

* Extend.
* Release.
* Inspect.

---

# 15. Waitlist

List:

* Customer.
* Departure.
* Travelers.
* Position.
* Joined At.
* Status.

Actions:

* Offer Slot.
* Reserve Slot.
* Confirm.
* Expire Offer.
* Remove.

Statuses:

```text
WAITING
OFFERED
RESERVED
CONFIRMED
EXPIRED
REMOVED
```

---

# 16. Pickup Points

---

## Fields

* Name *
* Address *
* Latitude.
* Longitude.
* Map Picker.
* Instructions.
* Contact Phone.
* Active.
* Display Order.

Actions:

* Add.
* Edit.
* Disable.
* Archive.

---

# 17. Payments

---

# 17.1 Payments List

Columns:

* Payment Number.
* Customer.
* Order.
* Type.
* Expected Amount.
* Paid.
* Remaining.
* Method.
* Status.
* Reference.
* Created.
* Reviewer.

Types:

```text
ORDER
WALLET_TOPUP
ORDER_BALANCE_DUE
```

---

# 17.2 Payment Detail

Sections:

* Payment Summary.
* Customer.
* Order.
* Bank Account Snapshot.
* Proofs.
* Bank Transactions.
* Allocations.
* Timeline.
* Internal Notes.
* Audit.

Actions بحسب الحالة:

* Review Proof.
* Request Information.
* Reject Proof.
* Match.
* Partial Match.
* Approve.
* Put On Hold.
* Reverse.

---

# 18. Payment Proofs

List:

* Proof ID.
* Customer.
* Payment.
* Amount Declared.
* Date.
* Sender.
* File.
* Duplicate Flag.
* Status.
* Submitted At.

Actions:

* Open File.
* Approve Proof.
* Reject.
* Request Better Copy.
* Mark Suspicious.

---

# 19. Bank Accounts

المسار:

```text
/settings/payments/bank-accounts
```

---

## Add Bank Account

حقول البنك فقط:

* رقم الحساب *
* اسم صاحب الحساب *
* اسم البنك *
* الفرع *

System Fields:

* Active.
* Default.
* Display Order.

لا يوجد:

* IBAN.
* SWIFT.
* BIC.

---

## Actions

* Add.
* Edit.
* Set Default.
* Enable.
* Disable.
* Archive.

تعديل الحساب لا يغير Payment Requests القديمة.

---

# 20. Bank Transactions

Finance يسجل الحركات التي يراها فعليًا في الحساب.

Fields:

* Bank Account.
* Amount.
* Transaction Date.
* Bank Reference.
* Sender Name.
* Description.
* Internal Notes.

Actions:

* Create.
* Edit قبل المطابقة.
* Mark To Check.
* Match.
* Partial Match.
* Mark Duplicate.
* Ignore.
* Reverse.

---

# 21. Reconciliation

Split view:

```text
Left:
Bank Transactions

Right:
Payment Requests / Wallet Top-ups
```

Go يقترح Matches بواسطة:

* Amount.
* Reference.
* Customer.
* Date.
* Sender.

لكن الموظف هو من يضغط:

```text
Confirm Match
```

---

# 22. Wallets

---

# 22.1 Wallet List

Columns:

* Customer.
* Wallet Number.
* Available.
* Reserved.
* Pending.
* Promotional.
* Status.
* Last Transaction.

Filters:

* Active.
* Frozen.
* Has Balance.
* Pending Top-up.

---

# 22.2 Wallet Detail

Header:

* Customer.
* Available.
* Reserved.
* Pending.
* Promotional.
* Status.

Tabs:

```text
Transactions
Top-ups
Holds
Orders
Refunds
Credits
Adjustments
Timeline
Audit
```

Actions:

* Freeze.
* Unfreeze.
* Create Adjustment.
* Add Credit.

لا يوجد:

```text
Set Balance
```

---

# 23. Wallet Top-ups

---

## User Flow

```text
Add Money
→ Amount
→ Bank Details
→ Transfer
→ Upload Proof
→ Review
→ Bank Match
→ Approval
→ Wallet Credit
```

---

## Admin List

Columns:

* Top-up Number.
* Customer.
* Amount.
* Reference.
* Bank.
* Proof Status.
* Top-up Status.
* Created.

Actions:

* Review.
* Approve.
* Reject.
* Request Info.
* Expire.
* Cancel.

---

# 24. Wallet Transactions

Types:

* Bank Top-up.
* Order Payment.
* Refund.
* Promotional Credit.
* Compensation.
* Adjustment.
* Reversal.

Columns:

* Transaction Number.
* Customer.
* Type.
* Amount.
* Direction.
* Balance After.
* Reference.
* Date.

لا يسمح Edit/Delete.

---

# 25. Wallet Adjustments

Add Adjustment:

* Wallet.
* Type:

  * Credit.
  * Debit.
* Amount.
* Reason.
* Reference.
* Internal Note.

قد يحتاج Approval حسب Threshold.

---

# 26. Wallet Holds

يعرض:

* Customer.
* Order.
* Amount.
* Created.
* Expires.
* Status.

Actions:

* Release عند الحالات الإدارية المناسبة.
* Inspect.

لا يسمح Capture يدويًا إلا Permission خاصة.

---

# 27. Refunds

---

## Refund List

Columns:

* Refund Number.
* Customer.
* Order.
* Amount.
* Method.
* Status.
* Created.
* Approved By.

---

## Create Refund

Fields:

* Order.
* Refundable Amount.
* Requested Amount.
* Reason.
* Internal Note.

Method:

```text
Wallet
Bank
Split
```

Split:

```text
Wallet Amount
Bank Amount
```

---

## Bank Refund Completion

Finance يدخل:

* Transfer Date.
* Bank Reference.
* Amount.
* Proof.
* Notes.

ثم:

```text
Complete Refund
```

---

# 28. Financial Ledger

Read-only operationally.

Pages:

```text
Accounts
Journal Entries
Journal Lines
```

Search by:

* Order.
* Payment.
* Refund.
* Wallet Transaction.
* Customer.
* Journal Number.

لا Edit/Delete.

Correction عبر Reversal Entry.

---

# 29. Pricing

قسم مستقل لإدارة قواعد التسعير.

---

## Pricing Rules List

* Name.
* Type.
* Applies To.
* Priority.
* Start.
* End.
* Status.

Types:

* Seasonal.
* Departure.
* Customer Segment.
* Volume.
* Manual.
* Package Variant.

---

## Add Pricing Rule

* Name.
* Description.
* Priority.
* Start At.
* End At.
* Target:

  * Packages.
  * Categories.
  * Destinations.
  * Departures.
* Calculation:

  * Fixed Price.
  * Increase %.
  * Decrease %.
  * Fixed Adjustment.
* Conditions.

---

# 30. Promotions

List:

* Name.
* Type.
* Discount.
* Target.
* Uses.
* Start.
* End.
* Status.

Add Promotion fields:

* Name.
* Description.
* Promotion Type.
* Discount Type.
* Discount Value.
* Applies To.
* Customer Segments.
* Minimum Order.
* Minimum Travelers.
* Max Uses.
* Max Per Customer.
* Start.
* End.
* Stack Policy.

---

# 31. Coupons

Fields:

* Code.
* Name.
* Description.
* Discount.
* Start.
* End.
* Usage Limit.
* Usage Per Customer.
* Eligible Packages.
* Eligible Segments.
* Minimum Spend.
* Active.

Actions:

* Generate.
* Bulk Generate.
* Disable.
* Export.
* View Usage.

---

# 32. Customer Segments

Builder:

```text
Condition
AND / OR
Condition
```

Available fields:

* Spend.
* Orders.
* Last Order.
* Location.
* Package Type.
* Wallet Balance.
* Cancellation Count.
* Customer Age.
* Registration Date.

Preview:

```text
Customers matching: 1,245
```

---

# 33. Loyalty

يجب أن يكون منفصلًا عن Wallet.

Settings:

* Enabled.
* Points earning.
* Points expiration.
* Tier rules.

Tiers:

* Bronze.
* Silver.
* Gold.
* VIP.

Rewards:

* Discount.
* Coupon.
* Upgrade.
* Add-on.

---

# 34. Referrals

Fields:

* Referral Code.
* Referrer.
* Referred User.
* Qualification Status.
* Reward Status.

Settings:

* Reward Referrer.
* Reward New Customer.
* Minimum qualifying order.
* Expiration.

---

# 35. Support

---

## Inbox

Columns:

* Conversation.
* Customer.
* Order.
* Assigned Agent.
* Priority.
* Status.
* Last Message.
* SLA.

---

## Conversation

* Customer Profile.
* Related Order.
* Messages.
* Attachments.
* Internal Notes.
* Tags.
* Assignment.

Actions:

* Reply.
* Add Internal Note.
* Assign.
* Escalate.
* Close.

---

# 36. Tickets

Fields عند الإنشاء:

* Customer.
* Subject.
* Category.
* Related Order.
* Priority.
* Description.
* Attachment.
* Assigned Team.

Statuses:

```text
OPEN
IN_PROGRESS
WAITING_CUSTOMER
WAITING_INTERNAL
ESCALATED
RESOLVED
CLOSED
```

---

# 37. Reviews

List:

* Customer.
* Package.
* Rating.
* Review.
* Verified Traveler.
* Status.
* Created.

Actions:

* Publish.
* Reject.
* Flag.
* Reply.
* Hide.

---

# 38. Vendors

إعادة بناء كاملة عن النظام الحالي.

Vendor Fields:

* Vendor Name.
* Legal Name.
* Email.
* Phone.
* Contact Person.
* Address.
* Services.
* Destinations.
* Status.
* Notes.

Vendor User يدخل عبر نظام Auth العادي + Role.

لا توجد Password داخل Vendor Collection.

---

# 39. Suppliers

Fields:

* Supplier Name.
* Contact.
* Phone.
* Email.
* Address.
* Service Types.
* Destinations.
* Payment Terms.
* Notes.
* Status.

Tabs:

```text
Services
Contracts
Departures
Orders
Payments
Documents
Notes
```

---

# 40. Contracts

Fields:

* Contract Number.
* Supplier.
* Name.
* Start Date.
* End Date.
* Currency.
* Terms.
* File.
* Status.
* Notes.

---

# 41. Content — Home Builder

الإدارة تتحكم في الصفحة الرئيسية للتطبيق بدون تحديث Flutter.

Sections:

* Hero.
* Banner.
* Featured Packages.
* Popular Destinations.
* Promotions.
* Categories.
* Custom Text.
* FAQ.

لكل Section:

* Enabled.
* Title.
* Subtitle.
* Data Source.
* Order.
* Start.
* End.
* Audience Segment.

---

# 42. Banners

Fields:

* Title.
* Subtitle.
* Image.
* Mobile Image.
* Action Type:

  * Package.
  * Destination.
  * URL.
  * Promotion.
  * None.
* Action Target.
* Start.
* End.
* Active.
* Display Order.

---

# 43. Pages

مثل:

* About.
* Contact.
* Terms.
* Privacy.
* Refund Policy.

Fields:

* Title.
* Slug.
* Content.
* SEO.
* Status.
* Publish Date.

---

# 44. FAQs

Fields:

* Question.
* Answer.
* Category.
* Order.
* Active.

CRUD + reorder.

---

# 45. Media Library

List/Grid.

Filters:

* Image.
* PDF.
* Video.
* Public.
* Private.

File detail:

* Filename.
* MIME.
* Size.
* Dimensions.
* Uploaded By.
* Created.
* Usage.
* Alt Text.

Actions:

* Rename metadata.
* Replace where safe.
* Archive.
* Delete if unused.

---

# 46. Notifications Center

هذا القسم مختلف عن Settings الخاصة بالإشعارات.

---

## Notification History

يعرض:

* Recipient.
* Type.
* Channel.
* Template.
* Status.
* Sent At.
* Delivered At.
* Failed Reason.

---

## Manual Notification

Admin يستطيع إرسال:

* لمستخدم واحد.
* Segment.
* مجموعة مستخدمين.
* جميع المستخدمين.

لكن يحتاج Permission خاصة.

Fields:

* Title.
* Body.
* Image.
* Deep Link.
* Audience.
* Schedule.
* Channel.

---

# 47. Analytics

كل صفحة يجب دعم:

* Date Range.
* Compare Previous Period.
* Filters.
* Export.

---

## Sales

* Gross.
* Net.
* Discounts.
* Refunds.
* AOV.
* Orders.
* Revenue by Package.
* Revenue by Destination.

---

## Customers

* New.
* Returning.
* Retention.
* LTV.
* Repeat Rate.
* Cohorts.

---

## Packages

* Views.
* Orders.
* Conversion.
* Revenue.
* Cancellation.
* Rating.
* Margin.

---

## Finance

* Received.
* Refunds.
* Net.
* Unmatched.
* Wallet Liability.
* Write-offs.

---

# 48. Custom Reports

Report Builder:

* Dataset.
* Columns.
* Filters.
* Grouping.
* Sorting.
* Visualization.

Actions:

* Run.
* Save.
* Duplicate.
* Schedule.
* Export.

---

# 49. Automation Workflows

Visual Builder:

```text
Trigger
   ↓
Condition
   ↓
Branch
   ↓
Action
```

---

## Triggers

* Order Created.
* Payment Approved.
* Proof Rejected.
* Wallet Top-up Approved.
* Departure Nearly Full.
* Document Expiring.
* Refund Requested.
* Support Ticket Created.

---

## Actions

* Send Notification.
* Send Email.
* Add Tag.
* Assign Staff.
* Create Task.
* Change Priority.
* Create Approval.
* Call Webhook.

لا يسمح Workflow باعتماد Payment تلقائيًا.

---

# 50. Approvals

Central Approval Inbox.

أنواع:

* Payment.
* Refund.
* Wallet Adjustment.
* Write-off.
* High Discount.
* Bank Account Change.
* High-value Order Edit.

Columns:

* Request.
* Type.
* Amount/Impact.
* Requested By.
* Created.
* Status.

Actions:

* Approve.
* Reject.
* Request Information.

---

# 51. Staff

List:

* Name.
* Email.
* Role.
* Team.
* Status.
* MFA.
* Last Login.

Add Staff:

* Name.
* Email.
* Role.
* Teams.
* Language.
* Require MFA.
* Send Invitation.

Actions:

* Edit.
* Suspend.
* Activate.
* Reset MFA.
* Revoke Sessions.

---

# 52. Roles

Role fields:

* Name.
* Description.
* Permissions.
* Teams.
* Active.

يمكن Duplicate Role.

---

# 53. Permissions

مجموعات:

```text
Orders
Customers
Catalog
Travel
Finance
Wallet
Refunds
Support
Content
Marketing
Analytics
Staff
System
```

كل Resource:

```text
Read
Create
Update
Delete/Archive
Approve
Export
Special Actions
```

---

# 54. Audit Logs

Filters:

* User.
* Action.
* Resource.
* Date.
* IP.

Record:

* Actor.
* Action.
* Resource.
* Before.
* After.
* Reason.
* IP.
* Device.
* Timestamp.

Audit لا يحذف.

---

# 55. Risk Flags

List:

* Resource.
* Flag.
* Severity.
* Detected At.
* Status.
* Assigned.

Examples:

* Duplicate Receipt.
* Duplicate Bank Reference.
* Suspicious Wallet Activity.
* Large Adjustment.
* Frequent Refunds.
* Multiple Failed Login.

Actions:

* Review.
* Resolve.
* Escalate.
* Add Note.

---

# 56. Imports

Supported:

* Customers.
* Packages.
* Products.
* Departures.
* Suppliers.
* Bank Statements.

Flow:

```text
Upload
→ Map Columns
→ Validate
→ Preview
→ Import
```

يظهر:

* Successful Rows.
* Failed Rows.
* Download Errors.

---

# 57. Exports

Can export:

* Orders.
* Customers.
* Travelers.
* Wallet.
* Payments.
* Ledger.
* Packages.
* Reports.

Format:

* CSV.
* XLSX.
* PDF where relevant.

Large exports تتم Background.

---

# 58. Integrations

كل Integration له صفحة:

* Name.
* Provider.
* Status.
* Enabled.
* Configuration.
* Last Success.
* Last Error.
* Test Connection.

Examples:

* Google OAuth.
* Maps.
* Weather.
* Email.
* Push.
* Object Storage.
* Analytics.
* AI/OCR.
* Accounting.

---

# 59. Webhooks

Fields:

* Name.
* URL.
* Signing Secret.
* Events.
* Active.
* Timeout.

Delivery Logs:

* Event.
* HTTP Status.
* Attempt.
* Response.
* Time.

Actions:

* Retry.
* Replay.
* Disable.

---

# 60. Feature Flags

Fields:

* Key.
* Description.
* Enabled.
* Environment.
* Rollout.
* Target Segment.

Examples:

```text
wallet_enabled
loyalty_enabled
waitlist_enabled
reviews_enabled
recommendations_enabled
```

---

# 61. Settings — أهم قسم يجب توسيعه

هذا القسم يجب ألا يكون مجرد App Constants.

المسار:

```text
/settings
```

يضم عدة مجموعات مستقلة.

---

# 62. Settings → General

Fields:

### Platform

* Platform Name.
* Admin Panel Name.
* Short Name.
* Company Name.
* Company Legal Name.
* Default Language.
* Default Timezone.
* Default Currency.
* Support Email.
* Support Phone.

### Admin Interface

* اسم لوحة التحكم.
* Browser Title.
* Sidebar Logo.
* Login Page Logo.
* Favicon.
* Compact Logo.
* Dark Logo.
* Light Logo.

### Footer

* Footer Text.
* Copyright Text.
* Company Website.

---

# 63. Settings → Branding

Fields:

* Primary Logo.
* Secondary Logo.
* Dark Logo.
* Light Logo.
* App Icon.
* Favicon.
* Splash Screen Logo.
* Email Logo.

Colors:

* Primary Color.
* Secondary Color.
* Accent Color.
* Success.
* Warning.
* Error.

Application:

* App Name.
* Short App Name.
* Android Display Name.
* iOS Display Name.

يمكن Preview قبل الحفظ.

---

# 64. Settings → Mobile App

هذه الإعدادات يقرأها Flutter من Backend.

Fields:

### General

* App Name.
* App Maintenance Mode.
* Maintenance Message.
* Customer Registration Enabled.
* Google Login Enabled.

### Versions

* Minimum Android Version.

* Latest Android Version.

* Force Android Update.

* Android Store URL.

* Minimum iOS Version.

* Latest iOS Version.

* Force iOS Update.

* iOS Store URL.

### Features

On/Off:

* Wallet.
* Reviews.
* Chat.
* Recommendations.
* Waitlist.
* Loyalty.
* News إذا أعيد تفعيلها.
* Maps.
* 360 View.

### Customer Support

* Support Phone.
* Support Email.
* WhatsApp Number اختياري.
* Working Hours.

---

# 65. Settings → Push Notifications

قسم مهم جدًا.

### Global

* Push Notifications Enabled.
* Default Notification Icon.
* Default Notification Sound.
* Default Image.
* Notification TTL.
* Enable Badge Counter.

### Providers

يخزن Backend Credentials بأمان وليس Frontend.

مثلاً:

* Android Push Provider.
* iOS APNs.

### APNs

* Team ID.
* Key ID.
* Bundle ID.
* Private Key Secret Reference.
* Environment:

  * Sandbox.
  * Production.

لا يتم عرض Private Key بعد حفظه.

### Android

إعداد Provider الذي سيستخدمه النظام للـPush.

Credentials داخل Secret Manager.

### Test

زر:

```text
Send Test Notification
```

مع:

* Select User.
* Device.
* Title.
* Body.

---

# 66. Settings → Notification Templates

لكل Event:

```text
ORDER_CREATED
PAYMENT_REQUIRED
PAYMENT_APPROVED
PAYMENT_REJECTED
WALLET_TOPUP_APPROVED
REFUND_COMPLETED
TRIP_REMINDER
DOCUMENT_REQUIRED
```

Template يحتوي:

### Arabic

* Title.
* Body.

### English

* Title.
* Body.

### Channels

* Push.
* In-app.
* Email.

### Deep Link

مثل:

```text
/orders/{{order_id}}
```

### Variables

مثل:

```text
{{customer_name}}
{{order_number}}
{{amount}}
{{departure_date}}
```

---

# 67. Settings → Notification Scheduling

إعدادات:

* Trip Reminder قبل:

  * X days.
  * X hours.

* Payment Reminder بعد:

  * X hours.

* Document Expiry Reminder:

  * 90 days.
  * 30 days.
  * 7 days.

* Wallet Top-up Reminder.

---

# 68. Settings → Email

Fields:

* Email Enabled.
* Provider.
* SMTP Host.
* SMTP Port.
* Username.
* Password Secret.
* Encryption.
* From Email.
* From Name.
* Reply-To.

Actions:

* Save.
* Test Email.

---

# 69. Settings → Payments

```text
Payments Enabled
Bank Transfer Enabled
Wallet Enabled
Mixed Payment Enabled
```

### Payment Request

* Default Expiration Hours.
* Allow Partial Payment.
* Allow Overpayment.
* Allow Underpayment.
* Require Proof.
* Allow Proof Replacement.

### Approval

* Require Manual Review.
* Require Second Approval.
* Second Approval Threshold.

---

# 70. Settings → Bank Accounts

كما سبق:

* رقم الحساب.
* اسم صاحب الحساب.
* اسم البنك.
* الفرع.

إضافة:

* Active.
* Default.
* Order.

---

# 71. Settings → Wallet

Fields:

### General

* Wallet Enabled.
* Top-up Enabled.
* Wallet Payments Enabled.
* Wallet Refunds Enabled.

### Limits

* Minimum Top-up.
* Maximum Top-up.
* Maximum Wallet Balance إذا قررت السياسة ذلك.
* Maximum Wallet Payment.

### Holds

* Hold Expiration Minutes.

### Credits

* Promotional Credits Enabled.
* Promotional Credit Expiration Policy.

### Approval

* Manual Adjustment Approval Threshold.
* Require Second Approval.

---

# 72. Settings → Orders

Fields:

* Order Prefix.
* Starting Sequence.
* Allow Draft Orders.
* Allow Admin Order Edit.
* Require Edit Reason.
* Allow Order Duplicate.
* Auto Archive Completed After X Days.

### Reservation

* Checkout Reservation Duration.
* Automatically Release Expired Reservations.

### Order Editing

* Allow Edit After Payment.
* Require Manager Approval Over Amount.
* Allow Package Change.
* Allow Departure Change.

---

# 73. Settings → Cancellation

يجب أن يكون Rule Builder.

مثال:

```text
Rule 1
More than 72h
Fee = 0%

Rule 2
48–72h
Fee = X%

Rule 3
24–48h
Fee = Y%

Rule 4
Less than 24h
Non-refundable
```

لكل Rule:

* From Hours.
* To Hours.
* Fee Type.
* Fee Value.
* Refund Percentage.
* Enabled.
* Priority.

---

# 74. Settings → Refunds

* Refunds Enabled.
* Wallet Refund Enabled.
* Bank Refund Enabled.
* Split Refund Enabled.
* Require Approval.
* Second Approval Threshold.
* Require Refund Proof.
* Require Reason.

---

# 75. Settings → Pricing

* Default Currency.
* Rounding Policy.
* Discount Stacking Policy:

  * Stack.
  * Best Deal.
  * Exclusive.
* Manual Discount Enabled.
* Manual Discount Max %.
* Manager Approval Threshold.

---

# 76. Settings → Travel

* Default Booking Lead Time.
* Default Capacity.
* Waitlist Default.
* Passport Minimum Validity.
* Document Reminder Days.
* Default Pickup Policy.

---

# 77. Settings → Customer

* Registration Enabled.
* Email Verification Required.
* Phone Required.
* Duplicate Email Policy.
* Duplicate Phone Policy.
* Allow Customer Account Deletion.
* Customer Deletion Review Required.

---

# 78. Settings → Security

### Authentication

* Password Minimum Length.
* Require Uppercase.
* Require Number.
* Require Special Character.
* Password Expiry إذا تم اعتماده.
* Failed Login Limit.
* Account Lock Duration.

### Sessions

* Customer Session TTL.
* Admin Session TTL.
* Refresh Token TTL.
* Maximum Active Admin Sessions.

### MFA

* Admin MFA Required.
* Finance MFA Required.
* Allowed MFA Methods.

### Network

* Admin IP Allowlist اختياري.
* Rate Limit Settings.

---

# 79. Settings → Files & Media

* Max Image Size.
* Max Document Size.
* Allowed Image Types.
* Allowed Document Types.
* Image Optimization Enabled.
* Thumbnail Generation.
* Private Signed URL TTL.
* Malware Scan Enabled.

---

# 80. Settings → Localization

* Languages.
* Default Language.
* RTL Enabled.
* Date Format.
* Time Format.
* Number Format.
* Currency Display.
* Timezone.

---

# 81. Settings → SEO

* Default Site Title.
* Default Meta Description.
* Default Social Image.
* App/Website URL.
* Robots Settings إذا وجد Web frontend.

---

# 82. Settings → Support

* Support Email.
* Support Phone.
* Support Working Hours.

SLA Defaults:

* Normal.
* High.
* Urgent.
* VIP.

---

# 83. Settings → Reviews

* Reviews Enabled.
* Require Completed Order.
* Auto Publish.
* Require Moderation.
* Allow Admin Replies.
* Minimum Rating.
* Maximum Rating.

---

# 84. Settings → Analytics

* Analytics Enabled.
* Data Retention.
* Event Tracking Enabled.
* Conversion Tracking Enabled.
* Internal Analytics Provider.

---

# 85. Settings → Integrations

Quick overview لكل Integration:

```text
Connected
Disconnected
Error
```

مع رابط لإعداداته.

---

# 86. Settings → Maintenance

* Customer App Maintenance.
* Admin Maintenance.
* Maintenance Message.
* Scheduled Start.
* Scheduled End.
* Allow Staff Bypass.

---

# 87. Settings → Feature Flags

واجهة مبسطة لتفعيل/إيقاف ميزات النظام.

---

# 88. Settings → Numbering & Prefixes

يمكن تخصيص:

```text
Order: ORD-
Payment: PAY-
Refund: REF-
Wallet Transaction: WLT-
Top-up: TOP-
Departure: DEP-
Customer: CUS-
```

مع Sequence مستقلة.

---

# 89. Settings → Data Retention

إعداد سياسات:

* Audit Retention.
* API Log Retention.
* Notification Log Retention.
* Temporary Upload Retention.
* Failed Job Retention.

السجلات المالية لا تطبق عليها نفس قواعد الحذف.

---

# 90. Settings → Backup & Recovery

عرض فقط من Admin العالي الصلاحية:

* Last Backup.
* Backup Status.
* Last Restore Test.
* Database PITR Enabled.
* Object Storage Backup Status.

لا يتم وضع Database Credentials هنا.

---

# 91. Go Backend — Modules النهائية

```text
internal/

auth/
identity/
sessions/
rbac/

customers/
segments/
travelers/
documents/

catalog/
categories/
destinations/
packages/
products/
addons/

pricing/
promotions/
coupons/

departures/
inventory/
reservations/
waitlist/
pickup/

checkout/
orders/
orderrevisions/
cancellations/

payments/
paymentproofs/
bankaccounts/
banktransactions/
reconciliation/

wallet/
wallettopups/
walletholds/
walletcredits/

refunds/
ledger/

support/
chat/
tickets/
reviews/

notifications/
push/
email/

vendors/
suppliers/
contracts/

cms/
media/

analytics/
reports/

automation/
approvals/
tasks/

staff/
audit/
risk/

integrations/
webhooks/
imports/
exports/

settings/
featureflags/

jobs/
outbox/

platform/
```

---

# 92. Settings Backend Architecture

لا نضع جميع الإعدادات في JSON ضخم واحد.

تقسيم:

```text
general_settings
branding_settings
mobile_app_settings
payment_settings
wallet_settings
order_settings
refund_settings
cancellation_settings
notification_settings
email_settings
security_settings
travel_settings
support_settings
media_settings
localization_settings
```

الإعدادات الحساسة مثل:

```text
SMTP password
APNs key
Storage credentials
OAuth secrets
```

لا تخزن كنص ظاهر.

يتم تخزين Secret Reference يشير إلى Secret Manager.

---

# 93. Settings Versioning

كل تعديل حساس في Settings يجب تسجيل:

```text
Before
After
Changed By
Reason
Changed At
```

خصوصًا:

* Bank Account.
* Payment Rules.
* Wallet Rules.
* Cancellation.
* Refund.
* Security.
* Push Credentials.

---

# 94. Cache Invalidation

بعض Settings تُقرأ باستمرار.

Go يمكنه Cache غير حساس، لكن عند تعديل Settings يجب إرسال Event:

```text
SettingsUpdated
```

لتحديث:

* API instances.
* Workers.
* Flutter Remote Config Endpoint.
* Next.js.

---

# 95. Flutter App Configuration Endpoint

مثلاً:

```text
GET /api/v1/app/config
```

يرجع فقط المعلومات العامة:

```text
app_name
maintenance
minimum_version
latest_version
force_update
support
enabled_features
branding
```

ولا يعيد أي Secrets.

---

# 96. Backend Business Rules

المبدأ الأساسي:

Next.js لا يقرر:

* سعر الطلب.
* حالة الدفع.
* Refund.
* Wallet Balance.
* Capacity.
* Discount النهائي.
* Approval.
* Authorization.

Go هو المسؤول.

---

# 97. State Machines

كل Domain حساس يجب أن يكون له State Machine.

مثل:

```text
Order
Payment
Refund
Wallet Top-up
Proof
Departure
Reservation
Traveler Document
Ticket
```

ولا يسمح Frontend بإرسال Status عشوائي.

بدل:

```text
PATCH status = "whatever"
```

يستخدم Commands:

```text
approve-payment
cancel-order
publish-package
freeze-wallet
```

---

# 98. Audit

أي عملية مثل:

```text
Edit Order
Approve Payment
Adjust Wallet
Change Bank Account
Change Refund Policy
Change Staff Role
Send Broadcast Notification
```

تدخل Audit Log.

---

# 99. Background Workers

Workers تتعامل مع:

* Push Notifications.
* Email.
* Scheduled Notifications.
* Reservation Expiration.
* Wallet Hold Expiration.
* Scheduled Publishing.
* Scheduled Unpublishing.
* Exports.
* Imports.
* Reports.
* Webhooks.
* Media Processing.
* Automation.
* Cleanup.

---

# 100. أهم مبدأ UI/UX

لوحة Next.js يجب أن تشبه الأنظمة التجارية الحديثة في سلوكها:

* Lists كثيفة لكن واضحة.
* Details pages مقسمة Cards/Tabs.
* Sticky action bar.
* Search دائم.
* Quick filters.
* Saved views.
* Bulk actions.
* Command palette.
* Side drawer لعرض تفاصيل سريعة.
* Confirmation dialogs للعمليات الحساسة.
* Reason field عند العمليات المهمة.
* Activity timeline.
* Unsaved changes warning.
* Optimistic locking.
* Permission-aware UI.

---

# 101. النتيجة النهائية

بعد هذه المواصفة لا يصبح لدينا مجرد:

```text
Products CRUD
Packages CRUD
Orders CRUD
Settings page
```

بل كل قسم عبارة عن Product Module مكتمل.

مثلاً Products لا يعني:

```text
Add Product
Edit Product
Delete Product
```

فقط.

بل:

```text
Product Catalog
→ Create
→ Classification
→ Pricing
→ Cost
→ Media
→ Availability
→ Package Eligibility
→ Purchase Rules
→ Refund Rules
→ Publishing
→ Scheduling
→ Versioned Changes
→ Analytics
→ Archive
```

والأمر نفسه ينطبق على:

```text
Packages
Orders
Customers
Travelers
Payments
Wallet
Refunds
Operations
Support
Marketing
Notifications
Staff
Settings
```

وقسم Settings يصبح **Control Plane للمنصة بأكملها**، ومنه يمكن إدارة:

```text
اسم المنصة
اسم لوحة التحكم
الشعار
الأيقونة
الألوان
اسم التطبيق
Maintenance Mode
إصدارات التطبيق
Force Update
ميزات التطبيق
بيانات الدعم
Push Notifications
Notification Templates
Email
Payment Rules
الحسابات البنكية
Wallet
Orders
Cancellation
Refund
Pricing
Travel
Security
Files
Languages
Reviews
Analytics
Integrations
Feature Flags
Numbering
Backups
```
