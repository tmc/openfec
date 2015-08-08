// Package openfec implements a client library for OpenFEC
//
// See https://api.open.fec.gov
//
// Overview
//
// Endpoints are classified in the following groups: Candidate, Committee, Financial, Filings and Schedules.
//
// Candidate
//
// Candidate endpoints give you access to information about the people running for office.
// This information is organized by candidate_id. If you're unfamiliar with candidate IDs,
// using `/candidates/search` will help you locate a particular candidate.
//
// Officially, a candidate is an individual seeking nomination for election to a federal
// office. People become candidates when they (or agents working on their behalf)
// raise contributions or make expenditures that exceed $5,000.
//
// The candidate endpoints primarily use data from FEC registration
// [Form 1](http://www.fec.gov/pdf/forms/fecfrm1.pdf), for candidate information, and
// [Form 2](http://www.fec.gov/pdf/forms/fecfrm2.pdf), for committee information.
//
//
// Committee
//
// Committees are entities that spend and raise money in an election. Their characteristics and
// relationships with candidates can change over time.
//
// You might want to use filters or search endpoints to find the committee you're looking
// for. Then you can use other committee endpoints to explore information about the committee
// that interests you.
//
// Financial information is organized by `committee_id`, so finding the committee you're interested in
// will lead you to more granular financial information.
//
// The committee endpoints include all FEC filers, even if they aren't registered as a committee.
//
// Officially, committees include the committees and organizations that file with the FEC.
// Several different types of organizations file financial reports with the FEC:
//
// * Campaign committees authorized by particular candidates to raise and spend funds in
// their campaigns
// * Non-party committees (e.g., PACs), some of which may be sponsored by corporations,
// unions, trade or membership groups, etc.
// * Political party committees at the national, state, and local levels
// * Groups and individuals making only independent expenditures
// * Corporations, unions, and other organizations making internal communications
//
// The committee endpoints primarily use data from FEC registration Form 1 and Form 2.
//
// Financial
//
// Fetch key information about a committee's Form 3, Form 3X, or Form 3P financial reports.
//
// Most committees are required to summarize their financial activity in each filing; those summaries
// are included in these files. Generally, committees file reports on a quarterly or monthly basis, but
// some must also submit a report 12 days before primary elections. Therefore, during the primary
// season, the period covered by this file may be different for different committees. These totals
// also incorporate any changes made by committees, if any report covering the period is amended.
//
// Information is made available on the API as soon as it's processed. Keep in mind, complex
// paper filings take longer to process.
//
// The financial endpoints use data from FEC [form 5](http://www.fec.gov/pdf/forms/fecfrm5.pdf),
// for independent expenditors; or the summary and detailed summary pages of the FEC
// [form 3](http://www.fec.gov/pdf/forms/fecfrm3.pdf), for House and Senate committees;
// [form 3X](http://www.fec.gov/pdf/forms/fecfrm3x.pdf), for PACs and parties;
// and [form 3P](http://www.fec.gov/pdf/forms/fecfrm3p.pdf), for presidential committees.
//
// Filings
//
// All official records and reports filed by or delivered to the FEC.
//
// Note: because the filings data includes many records, counts for large
// result sets are approximate.
//
// Schedules
//
// Schedules come from particular sections on forms and contain detailed transactional data.
//
// Schedule A explains where contributions come from. If you are interested in
// individual donors, this will be the endpoint you use.
//
// For the Schedule A aggregates, "memoed" items are not included to avoid double counting.
//
// Schedule B explains how money is spent.
package openfec
