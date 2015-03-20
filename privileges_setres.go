// +build dragonfly freebsd linux openbsd
// ^ mostly a guess, may need adjustment

package sprok

/*
 * Copyright 2015 Albert P. Tobey <atobey@datastax.com> @AlTobey
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import (
	"log"
	"syscall"
)

func (p *Process) drop_privileges() {
	var err error

	if p.Gid >= 0 {
		err = syscall.Setresgid(p.Gid, p.Gid, p.Gid)
		if err != nil {
			log.Fatalf("setresgid(%d...) failed: %s\n", p.Gid, err)
		}
	}

	if p.Uid >= 0 {
		err = syscall.Setresuid(p.Uid, p.Uid, p.Uid)
		if err != nil {
			log.Fatalf("setresuid(%d...) failed: %s\n", p.Uid, err)
		}
	}
}

