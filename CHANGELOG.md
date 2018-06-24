# Change Log

## [0.6.0](https://github.com/kubepack/onessl/tree/0.6.0) (2018-06-24)
[Full Changelog](https://github.com/kubepack/onessl/compare/0.5.0...0.6.0)

**Merged pull requests:**

- Fix prefix check for cert commands [\#43](https://github.com/kubepack/onessl/pull/43) ([tamalsaha](https://github.com/tamalsaha))
- Add peer-cert command [\#42](https://github.com/kubepack/onessl/pull/42) ([tamalsaha](https://github.com/tamalsaha))

## [0.5.0](https://github.com/kubepack/onessl/tree/0.5.0) (2018-06-12)
[Full Changelog](https://github.com/kubepack/onessl/compare/0.4.0...0.5.0)

**Merged pull requests:**

- Add support for hyperkube in doctor command [\#41](https://github.com/kubepack/onessl/pull/41) ([tamalsaha](https://github.com/tamalsaha))

## [0.4.0](https://github.com/kubepack/onessl/tree/0.4.0) (2018-06-07)
[Full Changelog](https://github.com/kubepack/onessl/compare/0.3.0...0.4.0)

**Closed issues:**

- doctor for kops cluster [\#35](https://github.com/kubepack/onessl/issues/35)
- Check kube-proxy running on master [\#33](https://github.com/kubepack/onessl/issues/33)
- check that SANS include `localhost` [\#32](https://github.com/kubepack/onessl/issues/32)
- Check if cert issued by ca.crt [\#30](https://github.com/kubepack/onessl/issues/30)

**Merged pull requests:**

- Capture proxy settings for kube-apiserver pods [\#40](https://github.com/kubepack/onessl/pull/40) ([tamalsaha](https://github.com/tamalsaha))
- Check if cert issued by ca.crt [\#39](https://github.com/kubepack/onessl/pull/39) ([tamalsaha](https://github.com/tamalsaha))
- Fix kube-apiserver detection for kops cluster [\#38](https://github.com/kubepack/onessl/pull/38) ([tamalsaha](https://github.com/tamalsaha))
- Check tls-cert-file for apiserver pods [\#37](https://github.com/kubepack/onessl/pull/37) ([tamalsaha](https://github.com/tamalsaha))
- Improve doctor command [\#36](https://github.com/kubepack/onessl/pull/36) ([tamalsaha](https://github.com/tamalsaha))
- Fix typo [\#31](https://github.com/kubepack/onessl/pull/31) ([tamalsaha](https://github.com/tamalsaha))
- Update changelog [\#29](https://github.com/kubepack/onessl/pull/29) ([tamalsaha](https://github.com/tamalsaha))

## [0.3.0](https://github.com/kubepack/onessl/tree/0.3.0) (2018-05-06)
[Full Changelog](https://github.com/kubepack/onessl/compare/0.2.0...0.3.0)

**Closed issues:**

- Check multi-master setup [\#19](https://github.com/kubepack/onessl/issues/19)
- Webhook Admission controller is enabled [\#14](https://github.com/kubepack/onessl/issues/14)

**Merged pull requests:**

- Fix doctor command for GKE [\#28](https://github.com/kubepack/onessl/pull/28) ([tamalsaha](https://github.com/tamalsaha))
- Update doctor command [\#27](https://github.com/kubepack/onessl/pull/27) ([tamalsaha](https://github.com/tamalsaha))
- Print cluster info at glog level 10 [\#26](https://github.com/kubepack/onessl/pull/26) ([tamalsaha](https://github.com/tamalsaha))
- Add doctor command [\#25](https://github.com/kubepack/onessl/pull/25) ([tamalsaha](https://github.com/tamalsaha))
- Update client-go to 7.0.0 [\#24](https://github.com/kubepack/onessl/pull/24) ([tamalsaha](https://github.com/tamalsaha))
- Error out if kube apiserver uses  insecure-skip-tls-verify : true [\#23](https://github.com/kubepack/onessl/pull/23) ([tamalsaha](https://github.com/tamalsaha))
- Add changelog [\#22](https://github.com/kubepack/onessl/pull/22) ([tamalsaha](https://github.com/tamalsaha))
- Support multiple orgs for client cert [\#21](https://github.com/kubepack/onessl/pull/21) ([tamalsaha](https://github.com/tamalsaha))
- Add backup cluster command [\#20](https://github.com/kubepack/onessl/pull/20) ([tamalsaha](https://github.com/tamalsaha))
- Add travis yaml [\#18](https://github.com/kubepack/onessl/pull/18) ([tahsinrahman](https://github.com/tahsinrahman))

## [0.2.0](https://github.com/kubepack/onessl/tree/0.2.0) (2018-03-10)
[Full Changelog](https://github.com/kubepack/onessl/compare/0.1.0...0.2.0)

**Merged pull requests:**

- Make onessl a kubectl plugin [\#17](https://github.com/kubepack/onessl/pull/17) ([tamalsaha](https://github.com/tamalsaha))

## [0.1.0](https://github.com/kubepack/onessl/tree/0.1.0) (2018-03-05)
**Fixed bugs:**

- onessl "Killed: 9" [\#9](https://github.com/kubepack/onessl/issues/9)
- MacOS version does not work [\#8](https://github.com/kubepack/onessl/issues/8)

**Closed issues:**

- has-key cfgmap|secret [\#13](https://github.com/kubepack/onessl/issues/13)
- wait-until-ready \<resource\> [\#12](https://github.com/kubepack/onessl/issues/12)
- envsubst [\#3](https://github.com/kubepack/onessl/issues/3)

**Merged pull requests:**

- Set analytics id properly [\#16](https://github.com/kubepack/onessl/pull/16) ([tamalsaha](https://github.com/tamalsaha))
- Add wait until ready commands [\#15](https://github.com/kubepack/onessl/pull/15) ([tamalsaha](https://github.com/tamalsaha))
- Rename --base flag to --minor [\#11](https://github.com/kubepack/onessl/pull/11) ([tamalsaha](https://github.com/tamalsaha))
- Use github.com/pkg/errors [\#10](https://github.com/kubepack/onessl/pull/10) ([tamalsaha](https://github.com/tamalsaha))
- Add envsubst command [\#7](https://github.com/kubepack/onessl/pull/7) ([tamalsaha](https://github.com/tamalsaha))
- Add semver command [\#6](https://github.com/kubepack/onessl/pull/6) ([tamalsaha](https://github.com/tamalsaha))
- Add jsonpath command [\#5](https://github.com/kubepack/onessl/pull/5) ([tamalsaha](https://github.com/tamalsaha))
- Add base64 encode/decode command [\#2](https://github.com/kubepack/onessl/pull/2) ([tamalsaha](https://github.com/tamalsaha))
- Add get kube-ca [\#1](https://github.com/kubepack/onessl/pull/1) ([tamalsaha](https://github.com/tamalsaha))



\* *This Change Log was automatically generated by [github_changelog_generator](https://github.com/skywinder/Github-Changelog-Generator)*